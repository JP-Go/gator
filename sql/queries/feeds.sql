-- name: AddFeed :one
INSERT INTO feeds (
    id,
    name, 
    url,
    user_id, 
    created_at,
    updated_at
) VALUES ( $1,$2,$3,$4,$5,$6 )
RETURNING *;

-- name: GetFeedsWithUserName :many 
SELECT feeds.*,users.name as user_name FROM feeds
INNER JOIN users 
    ON users.id = feeds.user_id;

-- name: FindFeedByURL :one 
SELECT * FROM feeds WHERE url = $1;

-- name: FindFeedByID :one 
SELECT * FROM feeds WHERE id = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET 
    last_fetched_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC
NULLS FIRST
LIMIT 1;
