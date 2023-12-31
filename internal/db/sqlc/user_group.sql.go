// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user_group.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUserGroup = `-- name: CreateUserGroup :one
INSERT INTO user_group (user_id,group_id) VALUES ($1,$2) RETURNING id, user_id, group_id, created_at
`

type CreateUserGroupParams struct {
	UserID  pgtype.UUID `json:"user_id"`
	GroupID uuid.UUID   `json:"group_id"`
}

func (q *Queries) CreateUserGroup(ctx context.Context, arg CreateUserGroupParams) (UserGroup, error) {
	row := q.db.QueryRow(ctx, createUserGroup, arg.UserID, arg.GroupID)
	var i UserGroup
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.GroupID,
		&i.CreatedAt,
	)
	return i, err
}
