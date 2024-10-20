-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comment (
    `id` VARCHAR(100) PRIMARY KEY,
    `user_id` VARCHAR(100),
    `post_id` VARCHAR(100),
    `content` TEXT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comment;
-- +goose StatementEnd
