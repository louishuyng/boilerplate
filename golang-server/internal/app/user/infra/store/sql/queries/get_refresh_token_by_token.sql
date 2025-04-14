-- name: GetRefreshTokenByToken :one
SELECT id, user_id, token, expired_at
FROM refresh_token
WHERE token = $1
LIMIT 1;