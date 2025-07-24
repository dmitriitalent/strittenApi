-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
	description    	VARCHAR(1000) NOT NULL,
	place          	VARCHAR(100) NOT NULL,
	date           	DATE NOT NULL,
	count          	INT NOT NULL,
	fundraising    	INT NOT NULL,
	user_id        	BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
