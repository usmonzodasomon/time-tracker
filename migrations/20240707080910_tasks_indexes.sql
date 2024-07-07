-- +goose Up
CREATE INDEX IF NOT EXISTS idx_tasks_user_id ON tasks (user_id);

-- +goose Down
DROP INDEX IF EXISTS idx_tasks_user_id;