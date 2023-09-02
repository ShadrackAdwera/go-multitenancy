// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username,email,tenant_id, password) 
VALUES ($1,$2,$3, $4) 
RETURNING id, username, email, tenant_id, password, password_changed_at, created_at
`

type CreateUserParams struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	TenantID uuid.UUID `json:"tenant_id"`
	Password string    `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.TenantID,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.TenantID,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, tenant_id, password, password_changed_at, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.TenantID,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, tenant_id, password, password_changed_at, created_at FROM users 
ORDER BY id
LIMIT $1 
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.TenantID,
			&i.Password,
			&i.PasswordChangedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  username = COALESCE($1,username),
  email = COALESCE($2,email),
  tenant_id = COALESCE($3,tenant_id),
  password = COALESCE($4,password)
WHERE id = $5
RETURNING id, username, email, tenant_id, password, password_changed_at, created_at
`

type UpdateUserParams struct {
	Username pgtype.Text `json:"username"`
	Email    pgtype.Text `json:"email"`
	TenantID pgtype.UUID `json:"tenant_id"`
	Password pgtype.Text `json:"password"`
	ID       uuid.UUID   `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.TenantID,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.TenantID,
		&i.Password,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}