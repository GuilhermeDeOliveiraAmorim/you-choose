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
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS actors (
    actor_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS movies (
    movie_id text,
    title text,
    synopsis text,
    imdb_rating text,
    votes int,
    you_choose_rating float,
    poster text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS movies_movie_lists (
    movie_id text,
    movie_list_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS actors (
    actor_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS actors_movies (
    actor_id text,
    movie_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS writers (
    writer_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS writers_movies (
    writer_id text,
    movie_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS directors (
    director_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS directors_movies (
    director_id text,
    movie_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS genres (
    genre_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS genres_movies (
    genre_id text,
    movie_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS tags (
    tag_id text,
    name text,
    picture text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS tags_movie_lists (
    tag_id text,
    movie_list_id text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);

CREATE TABLE IF NOT EXISTS files (
    file_id text,
    entity_id text,
    name text,
    size text,
    extension text,
    average_color text,
    is_deleted boolean,
    created_at text,
    updated_at text,
    deleted_at text
);