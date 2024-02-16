DROP TYPE IF EXISTS users_status_enum;

CREATE TYPE users_status_enum as enum (
    'active',
    'nonactive',
    'blocked',
    'suspended'
);