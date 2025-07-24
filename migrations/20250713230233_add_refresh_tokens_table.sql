-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh_tokens (
    id              SERIAL PRIMARY KEY,
    refresh_token   VARCHAR(512) NOT NULL UNIQUE,
    user_id         BIGINT NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
-- +goose StatementEnd
