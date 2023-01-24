CREATE TABLE IF NOT EXISTS choosers
(
    id text,
    first_name text,
    last_name text,
    username text,
    password text,
    picture text,
    is_deleted boolean,
    created_at date,
    updated_at date,
    deleted_at date
);

CREATE TABLE IF NOT EXISTS movie_lists
(
    id text,
    title text,
    description text,
    picture text,
    is_deleted boolean,
    created_at date,
    updated_at date,
    deleted_at date
);