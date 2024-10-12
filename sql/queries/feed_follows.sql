-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, user_id, feed_id;

-- name: GetFeedFollowByUserID :many
SELECT id, created_at, updated_at, user_id, feed_id
FROM feed_follows
WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE id = $1 AND user_id = $2;