-- +goose Up
-- +goose StatementBegin
CREATE TABLE additional_datas (
    id              SERIAL PRIMARY KEY,
	key             VARCHAR(200) NOT NULL,
    value           VARCHAR(200) NOT NULL,
    event_id        BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE additional_datas;
-- +goose StatementEnd
