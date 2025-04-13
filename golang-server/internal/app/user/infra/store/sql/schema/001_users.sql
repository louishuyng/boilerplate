-- +goose Up
Create table users (
  id uuid primary key,
  email varchar(255) not null,
  password varchar(255) not null,
  display_name varchar(255) not null
);

-- +goose Down
Drop table users;
