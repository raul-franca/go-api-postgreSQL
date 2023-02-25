CREATE TABLE IF NOT EXISTS todos(

    id          serial   primary key,
    name        varchar,
    description text,
    done        boolean,
    created_at timestamptz NOT NULL DEFAULT (now() )
);