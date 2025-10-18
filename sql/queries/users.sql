-- name: CreateUser :one
insert into users(id,createdAt,updatedAt,name,api_key)
values($1,$2,$3,$4,
  encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
select * from users where api_key=$1;
