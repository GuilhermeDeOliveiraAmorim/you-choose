CREATE TABLE IF NOT EXISTS choosers (
    id text,
    first_name text,
    last_name text,
    username text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS movie_lists (
    id text,
    title text,
    description text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS choosers_movie_lists (
    chooser_id text,
    movie_list_id text,
    created_at text,
    updated_at text,
    deleted_at text
);