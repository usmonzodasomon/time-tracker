-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     passport_serie INTEGER NOT NULL,
     passport_number INTEGER NOT NULL,
     name VARCHAR(255) NOT NULL,
     surname VARCHAR(255) NOT NULL,
     patronymic VARCHAR(255),
     address TEXT NOT NULL,
     is_deleted BOOLEAN DEFAULT FALSE,
     UNIQUE (passport_serie, passport_number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
