-- name: CreateUser :one
INSERT INTO users (id, email, password, display_name)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;
