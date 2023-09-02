-- name: CreatePermission :one
INSERT INTO permission (name,description,policy_id) VALUES ($1,$2,$3) RETURNING *;