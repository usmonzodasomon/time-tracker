-- +goose Up
CREATE INDEX IF NOT EXISTS idx_time_entries_task_id ON time_entries (task_id);
CREATE INDEX IF NOT EXISTS idx_time_entries_start_time ON time_entries (start_time);
CREATE INDEX IF NOT EXISTS idx_time_entries_end_time ON time_entries (end_time);

-- +goose Down
DROP INDEX IF EXISTS idx_time_entries_task_id;
DROP INDEX IF EXISTS idx_time_entries_start_time;
DROP INDEX IF EXISTS idx_time_entries_end_time;
