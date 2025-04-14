-- name: GetRefreshTokenByUserID :one
SELECT id, user_id, token, expired_at
FROM refresh_token
WHERE user_id = $1
LIMIT 1;
