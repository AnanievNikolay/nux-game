-- +goose Up
CREATE TABLE users (
    id CHAR(32) NOT NULL DEFAULT '' PRIMARY KEY,
    username VARCHAR(255) NOT NULL DEFAULT '',
    phone VARCHAR(32) NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(`username`,`phone`)
);

CREATE TABLE user_token(
    user_id CHAR(32) NOT NULL DEFAULT '',
    token CHAR(36) NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME NOT NULL,
    PRIMARY KEY (user_id, token)
);

CREATE INDEX idx_token ON user_token(token);

-- +goose Down
DROP TRIGGER trg_users_updated_at;

DROP TABLE users;

DROP TABLE user_token;