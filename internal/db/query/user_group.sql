-- name: CreateUserGroup :one
INSERT INTO user_group (user_id,group_id) VALUES ($1,$2) RETURNING *;