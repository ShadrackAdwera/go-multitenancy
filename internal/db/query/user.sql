-- name: CreateUser :one
INSERT INTO users (username,email,tenant_id, password) 
VALUES ($1,$2,$3, $4) 
RETURNING *;

-- name: ListUsers :many
SELECT * FROM user 
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: GetUser :one
SELECT * FROM user
WHERE id = $1;

-- name: UpdateUser :one
UPDATE user
SET
  username = COALESCE(sqlc.narg(username),username),
  email = COALESCE(sqlc.narg(email),email),
  tenant_id = COALESCE(sqlc.narg(tenant_id),tenant_id),
  password = COALESCE(sqlc.narg(password),password)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = $1;