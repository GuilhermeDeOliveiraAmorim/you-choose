-- Table: public.choosers

-- DROP TABLE IF EXISTS public.choosers;

CREATE TABLE IF NOT EXISTS public.choosers
(
    id text COLLATE pg_catalog."default",
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    username text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    picture text COLLATE pg_catalog."default",
    is_deleted boolean,
    created_at date,
    updated_at date,
    deleted_at date
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.choosers
    OWNER to root;