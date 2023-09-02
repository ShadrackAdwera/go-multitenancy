package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateUserArgs struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TenantUserTxInput struct {
	TenantData CreateTenantParams `json:"tenant"`
	UserData   CreateUserArgs     `json:"user"`
}

type TenantUserTxOutput struct {
	Message string `json:"message"`
	Tenant  Tenant `json:"tenant"`
	User    User   `json:"user"`
}

var appPermissions = []struct {
	name        string
	description string
}{
	{
		name:        "user_read",
		description: "View all users",
	},
	{
		name:        "user_create",
		description: "Invite new users",
	},
	{
		name:        "user_remove",
		description: "Delete any user",
	},
}

func createTenantDb(ctx context.Context, pgxpool *pgxpool.Pool, tenantName string) error {
	_, err := pgxpool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", tenantName))

	if err != nil {
		return err
	}
	return nil
}

func addTenantToPool(ctx context.Context, connectionString string, tenantDatabasePool map[string]*pgxpool.Pool, poolLock *sync.Mutex, tenantName string) error {
	poolLock.Lock()
	defer poolLock.Unlock()

	config, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		return fmt.Errorf("error parsing connection string: %w", err)
	}

	// Set any additional pool configuration options if needed
	config.MaxConns = 10
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return fmt.Errorf("error creating connection pool: %w", err)
	}

	tenantDatabasePool[tenantName] = pool
	return nil
}

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return fmt.Errorf("cannot create new migrate instance: %w", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run up migration: %w", err)
	}
	return nil
}

func seedData(ctx context.Context, q *Queries) (TenantGroup, error) {
	groupInfo, err := q.CreateTenantGroup(ctx, "Admins")

	if err != nil {
		return TenantGroup{}, fmt.Errorf("cannot seed new group: %w", err)
	}

	tp, err := q.CreateTenantPolicy(ctx, CreateTenantPolicyParams{
		Name: "Tenant Admin",
		GroupID: pgtype.UUID{
			Bytes: groupInfo.ID,
			Valid: true,
		},
	})

	if err != nil {
		return TenantGroup{}, fmt.Errorf("cannot seed permission: %w", err)
	}

	for _, apP := range appPermissions {
		_, err := q.CreatePermission(ctx, CreatePermissionParams{
			Name:        apP.name,
			Description: apP.description,
			PolicyID:    tp.ID,
		})

		if err != nil {
			return TenantGroup{}, fmt.Errorf("cannot seed permission: %w", err)
		}
	}

	return groupInfo, nil
}

func (s *Store) CreateTenantUserTx(ctx context.Context, input TenantUserTxInput) (TenantUserTxOutput, error) {
	var output TenantUserTxOutput
	err := s.execTx(ctx, func(q *Queries) error {

		// stripe payment first
		// create tenant db
		err := createTenantDb(ctx, s.pgxpool, input.TenantData.Name)
		if err != nil {
			return err
		}
		// add tenant to connection pool
		source := fmt.Sprintf("postgres://postgres:password@localhost:5431/%s?sslmode=disable", input.TenantData.Name)
		err = addTenantToPool(ctx, source, s.tenantDatabasePool, s.poolLock, input.TenantData.Name)
		if err != nil {
			return err
		}
		// establish connection to database using this tenant's connection pool

		// run migrations
		err = runDBMigration("file://db/migrations", source)
		if err != nil {
			return err
		}

		// seed data (roles + permissions)
		// seed 1 group - Admins - get id of group

		// create policy Tenant Admin - get id of tenant admin : name: Tenant Admin, group_id: from admins
		// seed all permissions : name, description, policy_id
		grp, err := seedData(ctx, q)
		if err != nil {
			return err
		}
		// create the user + assign tenant admin role

		t, err := q.CreateTenant(ctx, input.TenantData)

		if err != nil {
			return err
		}

		u, err := q.CreateUser(ctx, CreateUserParams{
			Username: input.UserData.Username,
			Email:    input.UserData.Email,
			TenantID: t.ID,
			Password: input.UserData.Password,
		})

		if err != nil {
			return err
		}

		_, err = q.CreateUserGroup(ctx, CreateUserGroupParams{
			UserID: pgtype.UUID{
				Bytes: u.ID,
				Valid: true,
			},
			GroupID: grp.ID,
		})

		if err != nil {
			return err
		}

		output.Tenant = t
		output.User = u

		return nil
	})
	return output, err
}
