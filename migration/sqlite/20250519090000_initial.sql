-- +goose Up
CREATE TABLE users (
    id CHAR(36) NOT NULL DEFAULT '' PRIMARY KEY,
    username VARCHAR(255) NOT NULL DEFAULT '',
    token VARCHAR(36) NOT NULL DEFAULT '',
    issued_at INTEGER NOT NULL DEFAULT 0,
    phone VARCHAR(255) NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Тригер для оновлення `updated_at` при зміні запису
CREATE TRIGGER trg_users_updated_at
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- +goose Down
DROP TRIGGER IF EXISTS trg_users_updated_at;
DROP TABLE IF EXISTS users;
