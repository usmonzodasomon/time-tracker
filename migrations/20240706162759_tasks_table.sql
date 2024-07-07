-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
   id SERIAL PRIMARY KEY,
   user_id INT NOT NULL,
   name VARCHAR(255) NOT NULL,
   description TEXT,
   FOREIGN KEY(user_id) REFERENCES users(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
