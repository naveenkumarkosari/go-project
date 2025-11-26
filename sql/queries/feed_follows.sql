-- name: CreateFeedFollows :one
INSERT INTO  feed_follows(id,createdAt,updatedAt,user_id,feed_id) 
VALUES ( $1,$2,$3,$4,$5 )
RETURNING *;

-- name: GetUserFeedFollows :many
Select * from feed_follows where user_id=$1;
