-- name: CreateUser :one
INSERT INTO users (id, name, api_key) 
VALUES ($1, $2, encode(sha256(random()::text::bytea), 'hex'))
RETURNING id, name, created_at, updated_at, api_key;

-- name: GetUserByAPIKey :one
SELECT id, name, created_at, updated_at, api_key
FROM users
WHERE api_key = $1;

