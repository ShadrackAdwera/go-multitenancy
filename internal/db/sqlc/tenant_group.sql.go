// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: tenant_group.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTenantGroup = `-- name: CreateTenantGroup :one
INSERT INTO tenant_group (name) VALUES ($1) RETURNING id, name, created_at
`

func (q *Queries) CreateTenantGroup(ctx context.Context, name string) (TenantGroup, error) {
	row := q.db.QueryRow(ctx, createTenantGroup, name)
	var i TenantGroup
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getTenantGroup = `-- name: GetTenantGroup :one
SELECT id, name, created_at FROM tenant_group WHERE id = $1
`

func (q *Queries) GetTenantGroup(ctx context.Context, id uuid.UUID) (TenantGroup, error) {
	row := q.db.QueryRow(ctx, getTenantGroup, id)
	var i TenantGroup
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
