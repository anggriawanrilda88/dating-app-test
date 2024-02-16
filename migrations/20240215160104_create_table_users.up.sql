CREATE TABLE users (
    id bigserial not null,
    email varchar(255) not null,
    password varchar(255) not null,
    status users_status_enum default 'nonactive' not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz null,
    constraint users_pkey primary key (id)
);

CREATE UNIQUE INDEX users_email_index ON users(email);

CREATE INDEX users_password_index ON users(password);

CREATE INDEX users_created_at_index ON users(created_at);

CREATE INDEX users_deleted_at_index ON users(deleted_at);