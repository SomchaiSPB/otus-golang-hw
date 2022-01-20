-- +goose Up
CREATE table events
(
    id          serial primary key,
    title       varchar,
    description text,
    duration    integer,
    user_id     integer,
    notify_in   integer,
    created_at  timestamp not null default now(),
    updated_at  timestamp
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table if exists events
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
