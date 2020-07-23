package campaigns

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/repos"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/campaigns"
	"github.com/sourcegraph/sourcegraph/internal/gitserver/protocol"
	"github.com/sourcegraph/sourcegraph/internal/vcs/git"
	"github.com/sourcegraph/sourcegraph/internal/workerutil"
	"github.com/sourcegraph/sourcegraph/schema"
)

type GitserverClient interface {
	CreateCommitFromPatch(ctx context.Context, req protocol.CreateCommitFromPatchRequest) (string, error)
}

type reconciler struct {
	gitserverClient GitserverClient
	sourcer         repos.Sourcer
	store           *Store
}

func (r *reconciler) HandlerFunc() workerutil.HandlerFunc {
	return func(ctx context.Context, tx workerutil.Store, record workerutil.Record) error {
		return r.process(ctx, r.store.With(tx), record.(*campaigns.Changeset))
	}
}

func (r *reconciler) process(ctx context.Context, tx *Store, ch *campaigns.Changeset) error {
	log15.Warn("Processing changeset", "changeset", ch.ID)

	spec, err := tx.GetChangesetSpecByID(ctx, ch.CurrentSpecID)
	if err != nil {
		return err
	}

	action, err := determineAction(ch, &spec.Spec)
	if err != nil {
		return err
	}

	switch action.name {
	case "publish":
		log15.Warn("Publishing", "changeset", ch.ID)
		return r.publishChangeset(ctx, tx, ch, spec)

	case "none":
		log15.Warn("No action", "changeset", ch.ID)

	default:
		return fmt.Errorf("TODO: action %q not implemented", action.name)
	}

	return nil
}

type reconcilerAction struct {
	name string
}

func determineAction(ch *campaigns.Changeset, desc *campaigns.ChangesetSpecDescription) (reconcilerAction, error) {
	var action reconcilerAction

	switch ch.PublicationState {
	case campaigns.ChangesetPublicationStateUnpublished:
		if desc.Published {
			action.name = "publish"
		} else {
			action.name = "none"
		}
	case campaigns.ChangesetPublicationStatePublished:
		ok, err := needsUpdate(ch, desc)
		if err != nil {
			return action, nil
		}
		if ok {
			action.name = "update"
		} else {
			action.name = "none"
		}
	// TODO: what about errored?
	default:
		return action, errors.New("implement updating already existing changeset")
	}

	return action, nil
}

func needsUpdate(ch *campaigns.Changeset, desc *campaigns.ChangesetSpecDescription) (bool, error) {
	currentTitle, err := ch.Title()
	if err != nil {
		return false, err
	}

	if currentTitle != desc.Title {
		return true, nil
	}

	currentBody, err := ch.Body()
	if err != nil {
		return false, err
	}

	if currentBody != desc.Body {
		return true, nil
	}

	currentBaseRef, err := ch.BaseRef()
	if err != nil {
		return false, err
	}

	if git.EnsureRefPrefix(currentBaseRef) != git.EnsureRefPrefix(desc.BaseRef) {
		return true, nil
	}

	// TODO: Where do we get the _current_ diff?
	// Do we get it from `gitserver`?
	//
	// But then we'd need to compare two strings, format-independent.
	//
	// Or do we ask gitserver to reapply the new diff and then compare
	// the current commit and the new one?

	return false, nil
}

func (r *reconciler) publishChangeset(ctx context.Context, tx *Store, ch *campaigns.Changeset, spec *campaigns.ChangesetSpec) (err error) {
	reposStore := repos.NewDBStore(tx.Handle().DB(), sql.TxOptions{})

	// Load associations from database

	campaignSpec, err := tx.GetCampaignSpec(ctx, GetCampaignSpecOpts{ID: spec.CampaignSpecID})
	if err != nil {
		return errors.Wrap(err, "failed to load campaign spec")
	}

	campaign, err := tx.GetCampaign(ctx, GetCampaignOpts{CampaignSpecID: campaignSpec.ID})
	if err != nil {
		return errors.Wrap(err, "failed to load campaign")
	}

	repo, err := loadRepo(ctx, reposStore, ch.RepoID)
	if err != nil {
		return errors.Wrap(err, "failed to load repository")
	}

	extSvc, err := loadExternalService(ctx, reposStore, repo)
	if err != nil {
		return errors.Wrap(err, "failed to load external service")
	}

	// Set up a source with which we can create a changeset
	ccs, err := r.buildChangesetSource(repo, extSvc)
	if err != nil {
		return err
	}

	// Create a commit and push it
	desc := spec.Spec
	if desc.IsExisting() {
		return errors.New("implement manual tracking of changesets")
	}

	diff, err := desc.Diff()
	if err != nil {
		return err
	}

	commitMessage, err := desc.CommitMessage()
	if err != nil {
		return err
	}

	opts := protocol.CreateCommitFromPatchRequest{
		Repo:       api.RepoName(repo.Name),
		BaseCommit: api.CommitID(desc.BaseRev),
		// IMPORTANT: We add a trailing newline here, otherwise `git apply`
		// will fail with "corrupt patch at line <N>" where N is the last line.
		Patch:     diff + "\n",
		TargetRef: desc.HeadRef,
		// TODO: if the changeset exists, this needs to be false, because we want to push the old branch
		UniqueRef: true,
		CommitInfo: protocol.PatchCommitInfo{
			Message:     commitMessage,
			AuthorName:  "THE INCREDIBLE SOURCEGRAPH RECONCILER",
			AuthorEmail: "campaigns@sourcegraph.com",
			Date:        spec.CreatedAt,
		},
		// We use unified diffs, not git diffs, which means they're missing the
		// `a/` and `/b` filename prefixes. `-p0` tells `git apply` to not
		// expect and strip prefixes.
		GitApplyArgs: []string{"-p0"},
		Push:         true,
	}
	ref, err := r.pushCommit(ctx, opts)
	if err != nil {
		return err
	}
	if ch.ExternalBranch != "" && ch.ExternalBranch != ref {
		return fmt.Errorf("ref %q doesn't match ChangesetJob's branch %q", ref, ch.ExternalBranch)
	}
	ch.ExternalBranch = ref

	baseRef := "refs/heads/master"
	if desc.BaseRef != "" {
		baseRef = desc.BaseRef
	}

	cs := repos.Changeset{
		Title:     campaign.Name,
		Body:      commitMessage,
		BaseRef:   baseRef,
		HeadRef:   git.EnsureRefPrefix(ref),
		Repo:      repo,
		Changeset: ch,
	}

	// TODO: If we're updating the changeset, there's a race condition here.
	// It's possible that `CreateChangeset` doesn't return the newest head ref
	// commit yet, because the API of the codehost doesn't return it yet.
	exists, err := ccs.CreateChangeset(ctx, &cs)
	if err != nil {
		return errors.Wrap(err, "creating changeset")
	}
	// If the Changeset already exists and our source can update it, we try to update it
	if exists {
		outdated, err := isOutdated(&cs)
		if err != nil {
			return errors.Wrap(err, "could not determine whether changeset needs update")
		}

		if outdated {
			if err := ccs.UpdateChangeset(ctx, &cs); err != nil {
				return errors.Wrap(err, "updating changeset")
			}
		}
	}

	events := ch.Events()
	SetDerivedState(ctx, ch, events)

	if err := tx.UpsertChangesetEvents(ctx, events...); err != nil {
		log15.Error("UpsertChangesetEvents", "err", err)
		return err
	}

	log15.Warn("Published changeset")

	ch.CreatedByCampaign = true
	ch.PublicationState = campaigns.ChangesetPublicationStatePublished
	return tx.UpdateChangeset(ctx, ch)
}

func (r *reconciler) pushCommit(ctx context.Context, opts protocol.CreateCommitFromPatchRequest) (string, error) {
	ref, err := r.gitserverClient.CreateCommitFromPatch(ctx, opts)
	if err != nil {
		if diffErr, ok := err.(*protocol.CreateCommitFromPatchError); ok {
			return "", errors.Errorf(
				"creating commit from patch for repository %q: %s\n"+
					"```\n"+
					"$ %s\n"+
					"%s\n"+
					"```",
				diffErr.RepositoryName, diffErr.InternalError, diffErr.Command, strings.TrimSpace(diffErr.CombinedOutput))
		}
		return "", err
	}

	return ref, nil
}

func (r *reconciler) buildChangesetSource(repo *repos.Repo, extSvc *repos.ExternalService) (repos.ChangesetSource, error) {
	sources, err := r.sourcer(extSvc)
	if err != nil {
		return nil, err
	}
	if len(sources) != 1 {
		return nil, errors.New("invalid number of sources for external service")
	}
	src := sources[0]
	ccs, ok := src.(repos.ChangesetSource)
	if !ok {
		return nil, errors.Errorf("creating changesets on code host of repo %q is not implemented", repo.Name)
	}

	return ccs, nil
}

func loadRepo(ctx context.Context, tx repos.Store, id api.RepoID) (*repos.Repo, error) {
	rs, err := tx.ListRepos(ctx, repos.StoreListReposArgs{IDs: []api.RepoID{id}})
	if err != nil {
		return nil, err
	}
	if len(rs) != 1 {
		return nil, errors.Errorf("repo not found: %d", id)
	}
	return rs[0], nil
}

func loadExternalService(ctx context.Context, reposStore repos.Store, repo *repos.Repo) (*repos.ExternalService, error) {
	var externalService *repos.ExternalService
	{
		args := repos.StoreListExternalServicesArgs{IDs: repo.ExternalServiceIDs()}

		es, err := reposStore.ListExternalServices(ctx, args)
		if err != nil {
			return nil, err
		}

		for _, e := range es {
			cfg, err := e.Configuration()
			if err != nil {
				return nil, err
			}

			switch cfg := cfg.(type) {
			case *schema.GitHubConnection:
				if cfg.Token != "" {
					externalService = e
				}
			case *schema.BitbucketServerConnection:
				if cfg.Token != "" {
					externalService = e
				}
			case *schema.GitLabConnection:
				if cfg.Token != "" {
					externalService = e
				}
			}
			if externalService != nil {
				break
			}
		}
	}

	if externalService == nil {
		return nil, errors.Errorf("no external services found for repo %q", repo.Name)
	}

	return externalService, nil
}

func isOutdated(c *repos.Changeset) (bool, error) {
	currentTitle, err := c.Changeset.Title()
	if err != nil {
		return false, err
	}

	if currentTitle != c.Title {
		return true, nil
	}

	currentBody, err := c.Changeset.Body()
	if err != nil {
		return false, err
	}

	if currentBody != c.Body {
		return true, nil
	}

	currentBaseRef, err := c.Changeset.BaseRef()
	if err != nil {
		return false, err
	}

	if git.EnsureRefPrefix(currentBaseRef) != git.EnsureRefPrefix(c.BaseRef) {
		return true, nil
	}

	return false, nil
}

// combineErrors returns a multierror containing all fo the non-nil error parameter values.
// This method should be used over multierror when it is not guaranteed that the original
// error was non-nil (multierror.Append creates a non-nil error even if it is empty).
func combineErrors(errs ...error) (err error) {
	for _, e := range errs {
		if e != nil {
			if err == nil {
				err = e
			} else {
				err = multierror.Append(err, e)
			}
		}
	}

	return err
}
