-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user (
    `id` VARCHAR(30) PRIMARY KEY,
    `email` VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    `display_name` VARCHAR(100) NOT NULL,
    `profile_picture_url` VARCHAR(100) NOT NULL,
    `user_type` ENUM('admin', 'user', 'member'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
-- +goose StatementEnd
