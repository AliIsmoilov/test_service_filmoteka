DROP TABLE IF EXISTS actors CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

DROP TYPE IF EXISTS gender;
CREATE TYPE gender AS ENUM('male', 'female');

CREATE TABLE actors
(
    id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name    VARCHAR(512)  NOT NULL CHECK ( name <> '' ),
    gender gender,
    birth_date VARCHAR(56),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);