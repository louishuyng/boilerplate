-- name: GetUserByEmail :one
SELECT id, email, password, display_name
FROM users
WHERE email = $1
LIMIT 1;