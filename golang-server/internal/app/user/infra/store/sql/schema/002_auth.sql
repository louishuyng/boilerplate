-- +goose Up
CREATE TABLE refresh_token (
  id uuid primary key,
  user_id uuid not null,
  token varchar(255) not null,
  expired_at timestamp not null,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE refresh_token;
