-- name: DeleteRefreshTokenByUserID :exec
DELETE FROM refresh_token
WHERE user_id = $1;