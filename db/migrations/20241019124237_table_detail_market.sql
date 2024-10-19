-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS detail_market (
    `id` VARCHAR(100) PRIMARY KEY,
    `post_id` VARCHAR(100) NOT NULL,
    `product_name` VARCHAR(100) NOT NULL,
    `product_type` VARCHAR(50) NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,
    `seller_address` VARCHAR(255) NOT NULL,
    `phone_number` VARCHAR(15) NOT NULL,
    `description` TEXT NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS detail_market;
-- +goose StatementEnd
