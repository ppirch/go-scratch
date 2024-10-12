-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING id, name, url, created_at, updated_at, user_id, last_fetched_at;

-- name: GetFeeds :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at
FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING id, name, url, created_at, updated_at, user_id, last_fetched_at;
