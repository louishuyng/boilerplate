-- name: UpdateRefreshTokenExpiredAt :exec
UPDATE refresh_token
SET expired_at = $2
WHERE token = $1;
