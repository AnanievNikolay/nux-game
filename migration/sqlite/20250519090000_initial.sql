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

CREATE INDEX idx_user_token ON user_token(token);

CREATE TABLE games (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id CHAR(36) NOT NULL,
    number INT NOT NULL DEFAULT 0,
    is_win BOOLEAN NOT NULL DEFAULT 0,
    prize REAL NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL
);

CREATE INDEX idx_game_user ON games(user_id);

-- +goose Down
DROP TRIGGER trg_users_updated_at;

DROP TABLE users;

DROP INDEX idx_user_token;

DROP TABLE user_token;

DROP INDEX idx_game_user;

DROP TABLE games;