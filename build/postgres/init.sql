DROP DATABASE myService;

CREATE DATABASE myService
    WITH
    OWNER = postgres
    CONNECTION LIMIT = -1
    ;

CREATE COLLATION posix (LOCALE = 'POSIX');
CREATE EXTENSION citext;


CREATE TABLE AdsPosts (
    id serial NOT NULL PRIMARY KEY,
    title text NOT NULL UNIQUE,
    description text NOT NULL,
    photos text[] NOT NULL,
    price bigint NOT NULL,
    date timestamp with time zone NOT NULL
);

CREATE INDEX AdsPosts_price ON AdsPosts (price);
CREATE INDEX AdsPosts_date ON AdsPosts (date);
