-- +migrate Down
DROP TABLE IF EXISTS users CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS merch CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS purchases CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS transactions CASCADE;

-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id       INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    coins    INTEGER DEFAULT 1000
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS merch (
    id    INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name  VARCHAR(255) NOT NULL UNIQUE,
    price INTEGER      NOT NULL CHECK (price > 0)
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS purchases (
    id         INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    merch_id   INTEGER REFERENCES merch (id) ON DELETE CASCADE
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS transactions (
    id           INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    from_user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    to_user_id   INTEGER REFERENCES users (id) ON DELETE CASCADE,
    amount       INTEGER NOT NULL CHECK (amount > 0),
    CONSTRAINT ck_transactions_not_equals_from_uid_to_uid CHECK (from_user_id <> to_user_id)
);