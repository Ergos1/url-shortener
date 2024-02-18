-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE urls(
    id BIGSERIAL PRIMARY KEY,
    short_url VARCHAR(255) UNIQUE NOT NULL,
    original_url TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE urls;
-- +goose StatementEnd
