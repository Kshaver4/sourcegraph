package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/repos"
	ee "github.com/sourcegraph/sourcegraph/enterprise/internal/campaigns"
	ct "github.com/sourcegraph/sourcegraph/enterprise/internal/campaigns/testing"
	"github.com/sourcegraph/sourcegraph/internal/actor"
	"github.com/sourcegraph/sourcegraph/internal/campaigns"
	"github.com/sourcegraph/sourcegraph/internal/db/dbutil"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/github"
)

func main() {
	ctx := context.Background()

	dsn := dbutil.PostgresDSN("sourcegraph", os.Getenv)
	db, err := dbutil.NewDB(dsn, "campaigns-reconciler")
	if err != nil {
		log.Fatalf("failed to initialize db store: %v", err)
	}

	if _, err := db.ExecContext(ctx, "DELETE FROM campaigns;"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "DELETE FROM changesets;"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "DELETE FROM changeset_specs;"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.ExecContext(ctx, "DELETE FROM campaign_specs;"); err != nil {
		log.Fatal(err)
	}

	campaignsStore := ee.NewStore(db)
	reposStore := repos.NewDBStore(db, sql.TxOptions{})

	repoName := "github.com/sourcegraph/automation-testing"
	rs, err := reposStore.ListRepos(ctx, repos.StoreListReposArgs{Names: []string{
		repoName,
	}})
	if err != nil {
		log.Fatal(err)
	}
	if len(rs) != 1 {
		log.Fatal(errors.Errorf("repo not found: %q", repoName))
	}
	repo := rs[0]

	campaignSpec, err := campaigns.NewCampaignSpecFromRaw(ct.TestRawCampaignSpec)
	if err != nil {
		log.Fatal(err)
	}
	campaignSpec.UserID = 1
	campaignSpec.NamespaceUserID = 1
	if err := campaignsStore.CreateCampaignSpec(ctx, campaignSpec); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created campaign spec.")

	repoID := graphqlbackend.MarshalRepositoryID(repo.ID)
	changesetSpec := &campaigns.ChangesetSpec{
		CampaignSpecID: campaignSpec.ID,
		UserID:         1,
		RepoID:         repo.ID,
		Spec: campaigns.ChangesetSpecDescription{
			BaseRepository: repoID,
			BaseRev:        "e4435274b43033cf0c212f61a2c16f7f2210cf56",
			BaseRef:        "refs/heads/master",
			HeadRepository: repoID,
			HeadRef:        "refs/heads/reconciler-test-3",
			Title:          "Testing the reconciler",
			Body:           "This should test the new reconciler",
			Commits: []campaigns.GitCommitDescription{
				{
					Message: "Testing the reconciler\n\nThis should test the new conciler",
					Diff: `diff --git test.md test.md
index 52ada66..0aaaf37 100644
--- test.md
+++ test.md
@@ -1 +1,3 @@
-# This is a test
+# This is a test.
+
+And this is another line
`,
				},
			},
			Published: true,
		},
	}

	if err := campaignsStore.CreateChangesetSpec(ctx, changesetSpec); err != nil {
		log.Fatalf("failed to create changeset spec: %s", err)
	}
	fmt.Println("Created changeset spec.")

	// The ApplyCampaign should create the changesets, but we do it "by hand" here

	fmt.Println("Applying campaign spec...")
	actorCtx := actor.WithActor(ctx, &actor.Actor{UID: 1, Internal: true})
	svc := ee.NewService(campaignsStore, nil)
	campaign, err := svc.ApplyCampaign(actorCtx, ee.ApplyCampaignOpts{CampaignSpecRandID: campaignSpec.RandID})
	if err != nil {
		log.Fatal(err)
	}

	ch := &campaigns.Changeset{
		RepoID:              changesetSpec.RepoID,
		ExternalID:          "TODO",
		ExternalServiceType: repo.ExternalRepo.ServiceType,

		CampaignIDs:   []int64{campaign.ID},
		CurrentSpecID: changesetSpec.ID,

		PublicationState: campaigns.ChangesetPublicationStateUnpublished,
		ReconcilerState:  "queued",
	}

	err = campaignsStore.CreateChangesets(ctx, ch)
	if err != nil {
		log.Fatalf("failed to create changesets: %s", err)
	}
	fmt.Println("done!")
}

func newGitHubPR() *github.PullRequest {
	githubActor := github.Actor{
		AvatarURL: "https://avatars2.githubusercontent.com/u/1185253",
		Login:     "mrnugget",
		URL:       "https://github.com/mrnugget",
	}

	return &github.PullRequest{
		ID:           "FOOBARID",
		Title:        "Fix a bunch of bugs",
		Body:         "This fixes a bunch of bugs",
		URL:          "https://github.com/sourcegraph/sourcegraph/pull/12345",
		Number:       12345,
		Author:       githubActor,
		Participants: []github.Actor{githubActor},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		HeadRefName:  "campaigns/test",
	}
}
