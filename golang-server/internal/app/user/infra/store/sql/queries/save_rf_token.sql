-- name: SaveRefreshToken :one
INSERT INTO refresh_token (id, user_id, token, expired_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;
