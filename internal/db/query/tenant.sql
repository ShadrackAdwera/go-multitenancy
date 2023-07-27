-- name: CreateTenant :one
INSERT INTO tenant (name,logo) VALUES ($1,$2) RETURNING *;

-- name: ListTenants :one
SELECT * FROM tenant ORDER BY id LIMIT $1 OFFSET $2;

-- name: GetTenant :one
SELECT * FROM tenant WHERE id = $1;

-- name: UpdateTenant :one
UPDATE tenant 
SET
  name = COALESCE(sqlc.narg(name),name),
  logo = COALESCE(sqlc.narg(logo),logo)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTenant :exec
DELETE FROM tenant 
WHERE id = $1;


