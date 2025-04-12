-- name: CreateExample :one
INSERT INTO example (id, name)
VALUES (
    gen_random_uuid(),
    $1
)
RETURNING *;
