-- name: CreateTenantPolicy :one
INSERT INTO tenant_policy (name,group_id) VALUES ($1,$2) RETURNING *;

-- name: GetTenantPolicy :one
SELECT * FROM tenant_policy WHERE id = $1;