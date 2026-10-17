-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM
    feed_follows
    Inner JOIN feeds ON feed_follows.feed_id = feeds.id
    Inner JOIN users ON feed_follows.user_id = users.id
WHERE
    feed_follows.user_id = $1;