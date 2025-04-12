-- +goose Up
Create table example (
  id uuid primary key,
	name varchar(255) not null
);

-- +goose Down
Drop table example;
