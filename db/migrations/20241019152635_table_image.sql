-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS post_image (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100),
    `file_path` VARCHAR(200),
    `image_order` INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS post_image;
-- +goose StatementEnd
