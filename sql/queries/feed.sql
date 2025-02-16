-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserFeeds :one
SELECT * FROM feed WHERE user_id = $1;

-- name GetAllFeeds
SELECT * FROM feed;

-- name GetOneFeed
SELECT * FROM feed WHERE id = $1;