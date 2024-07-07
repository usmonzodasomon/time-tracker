-- +goose Up
CREATE INDEX IF NOT EXISTS idx_users_is_deleted ON users (is_deleted);
CREATE INDEX IF NOT EXISTS idx_users_address ON users (address);
CREATE INDEX IF NOT EXISTS idx_users_name_surname_patronymic ON users (name, surname, patronymic);

-- +goose Down
DROP INDEX IF EXISTS idx_users_is_deleted;
DROP INDEX IF EXISTS idx_users_address;
DROP INDEX IF EXISTS idx_users_name_surname_patronymic;
