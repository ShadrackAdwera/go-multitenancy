-- name: CreateTenantGroup :one
INSERT INTO tenant_group (name) VALUES ($1) RETURNING *;

-- name: GetTenantGroup :one
SELECT * FROM tenant_group WHERE id = $1;