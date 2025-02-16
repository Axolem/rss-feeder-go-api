-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;