-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post (
    `id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100) NOT NULL,
    `post_type` VARCHAR(30) NOT NULL,
    `title` VARCHAR(100) NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post;
-- +goose StatementEnd
