BEGIN;

ALTER TABLE changesets RENAME COLUMN IF EXISTS current_spec_id TO changeset_spec_id;
ALTER TABLE changesets DROP COLUMN IF EXISTS previous_spec_id;

ALTER TABLE changesets DROP COLUMN IF EXISTS publication_state;
ALTER TABLE changesets ALTER COLUMN external_id SET NOT NULL;
ALTER TABLE changesets ALTER COLUMN metadata SET NOT NULL;
ALTER TABLE changesets ALTER COLUMN external_service_type SET NOT NULL;

ALTER TABLE changesets DROP COLUMN IF EXISTS reconciler_state;
ALTER TABLE changesets DROP COLUMN IF EXISTS failure_message;
ALTER TABLE changesets DROP COLUMN IF EXISTS started_at;
ALTER TABLE changesets DROP COLUMN IF EXISTS finished_at;
ALTER TABLE changesets DROP COLUMN IF EXISTS process_after;
ALTER TABLE changesets DROP COLUMN IF EXISTS num_resets;

COMMIT;
