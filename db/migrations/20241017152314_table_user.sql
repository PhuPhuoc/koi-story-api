-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user (
    `id` VARCHAR(100) PRIMARY KEY,
    `email` VARCHAR(30) NOT NULL,
    `password` VARCHAR(30) NOT NULL,
    `display_name` VARCHAR(100) NOT NULL,
    `profile_picture_url` VARCHAR(300) NOT NULL,
    `user_type` ENUM('admin', 'user', 'member'),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME
);

CREATE TABLE IF NOT EXISTS face_data (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` VARCHAR(100) NOT NULL,
    `img_url` LONGTEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS face_data;
-- +goose StatementEnd
