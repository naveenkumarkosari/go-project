-- name: CreateFeed :one
INSERT INTO  feeds(id,createdAt,updatedAt,content,createdBy) 
VALUES ( $1,$2,$3,$4,$5 )
RETURNING *;

-- name: GetUserPosts :one
SELECT * FROM feeds where createdBy=$1;
