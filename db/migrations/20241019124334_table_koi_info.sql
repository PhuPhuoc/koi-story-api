-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS koi_info (
    `id` VARCHAR(100) PRIMARY KEY,
    `detail_market_id` VARCHAR(100),
    `color` VARCHAR(100),
    `size` VARCHAR(100),
    `old` VARCHAR(100),
    `type` VARCHAR(100),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS koi_info;
-- +goose StatementEnd
