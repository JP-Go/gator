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

