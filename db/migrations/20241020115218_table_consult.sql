-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS detail_consult (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100),
    `content` TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS detail_consult;
-- +goose StatementEnd
