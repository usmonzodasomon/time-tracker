-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS time_entries (
    id SERIAL PRIMARY KEY,
    task_id INT NOT NULL,
    start_time TIMESTAMP DEFAULT now(),
    end_time TIMESTAMP,
    FOREIGN KEY(task_id) REFERENCES tasks(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS time_entries;
-- +goose StatementEnd
