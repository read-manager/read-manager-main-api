CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password bytea NOT NULL,
    nickname text NOT NULL,
    read_points integer NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NULL
);
