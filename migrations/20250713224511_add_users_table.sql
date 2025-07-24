-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    login           VARCHAR(30) NOT NULL UNIQUE,
    name            VARCHAR(50) NOT NULL,
    surname         VARCHAR(50) NOT NULL,
    email           VARCHAR(100) NOT NULL UNIQUE,
    password_hash   VARCHAR(100) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
