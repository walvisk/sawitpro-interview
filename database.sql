/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

CREATE TABLE users (
    id serial primary key,
    phone varchar(13) unique not null,
    country_code varchar(3) not null default '+62',
    password varchar(255) not null,
    full_name varchar(64) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz default null
);

CREATE INDEX idx_phone ON users (phone);

CREATE TABLE user_logs (
    id serial primary key,
    login_at timestamptz not null default current_timestamp,
    user_id integer references users(id) on delete cascade
);