package campaigns

import (
	"context"
	"database/sql"
	"time"

	"github.com/inconshreveable/log15"
	"github.com/keegancsmith/sqlf"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/repos"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"github.com/sourcegraph/sourcegraph/internal/metrics"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/internal/trace"
	"github.com/sourcegraph/sourcegraph/internal/workerutil"
)

// maxWorkers defines the maximum number of changeset jobs to run in parallel.
var maxWorkers = env.Get("CAMPAIGNS_MAX_WORKERS", "8", "maximum number of repository jobs to run in parallel")

const defaultWorkerCount = 8

func RunWorkers(ctx context.Context, s *Store, clock func() time.Time, gitClient GitserverClient, sourcer repos.Sourcer, backoffDuration time.Duration) {
	log15.Warn("RunWorkers")

	r := &reconciler{gitserverClient: gitClient, sourcer: sourcer, store: s}

	options := workerutil.WorkerOptions{
		Handler:     r.HandlerFunc(),
		NumHandlers: 1,
		Interval:    backoffDuration,

		Metrics: workerutil.WorkerMetrics{HandleOperation: newObservationOperation()},
	}

	workerStore := workerutil.NewStore(s.Handle(), workerutil.StoreOptions{
		TableName:            "changesets",
		AlternateColumnNames: map[string]string{"state": "reconciler_state"},
		ColumnExpressions: []*sqlf.Query{
			sqlf.Sprintf("changesets.id"),
			sqlf.Sprintf("changesets.repo_id"),
			sqlf.Sprintf("changesets.created_at"),
			sqlf.Sprintf("changesets.updated_at"),
			sqlf.Sprintf("changesets.metadata"),
			sqlf.Sprintf("changesets.campaign_ids"),
			sqlf.Sprintf("changesets.external_id"),
			sqlf.Sprintf("changesets.external_service_type"),
			sqlf.Sprintf("changesets.external_branch"),
			sqlf.Sprintf("changesets.external_deleted_at"),
			sqlf.Sprintf("changesets.external_updated_at"),
			sqlf.Sprintf("changesets.external_state"),
			sqlf.Sprintf("changesets.external_review_state"),
			sqlf.Sprintf("changesets.external_check_state"),
			sqlf.Sprintf("changesets.created_by_campaign"),
			sqlf.Sprintf("changesets.added_to_campaign"),
			sqlf.Sprintf("changesets.diff_stat_added"),
			sqlf.Sprintf("changesets.diff_stat_changed"),
			sqlf.Sprintf("changesets.diff_stat_deleted"),
			sqlf.Sprintf("changesets.sync_state"),
			sqlf.Sprintf("changesets.changeset_spec_id"),
			sqlf.Sprintf("changesets.publication_state"),
			sqlf.Sprintf("changesets.reconciler_state"),
			sqlf.Sprintf("changesets.failure_message"),
			sqlf.Sprintf("changesets.started_at"),
			sqlf.Sprintf("changesets.finished_at"),
			sqlf.Sprintf("changesets.process_after"),
			sqlf.Sprintf("changesets.num_resets"),
		},
		Scan:              scanFirstChangesetRecord,
		OrderByExpression: sqlf.Sprintf("changesets.updated_at"),
		StalledMaxAge:     5 * time.Second, // TODO(mrnugget)
		MaxNumResets:      5,               // TODO(mrnugget)
	})

	worker := workerutil.NewWorker(ctx, workerStore, options)
	worker.Start()
}

func scanFirstChangesetRecord(rows *sql.Rows, err error) (workerutil.Record, bool, error) {
	return scanFirstChangeset(rows, err)
}

func newObservationOperation() *observation.Operation {
	observationContext := &observation.Context{
		Logger:     log15.Root(),
		Tracer:     &trace.Tracer{Tracer: opentracing.GlobalTracer()},
		Registerer: prometheus.DefaultRegisterer,
	}

	metrics := metrics.NewOperationMetrics(
		observationContext.Registerer,
		"reconciler_todo",
		metrics.WithLabels("op"),
		metrics.WithCountHelp("Total number of results returned"),
	)

	return observationContext.Operation(observation.Op{
		Name:         "ReconcilerTODO.Process",
		MetricLabels: []string{"process"},
		Metrics:      metrics,
	})
}
