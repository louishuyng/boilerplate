-- name: UpdateRefreshTokenExpiredAt :exec
UPDATE refresh_token
SET expired_at = $2
WHERE id = $1;
