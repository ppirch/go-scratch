-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, url, created_at, updated_at, user_id;

-- name: GetFeeds :many
SELECT id, name, url, created_at, updated_at, user_id
FROM feeds;

