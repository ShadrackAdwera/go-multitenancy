package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ShadrackAdwera/go-utils/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomTenant(t *testing.T) Tenant {
	args := CreateTenantParams{
		Name: utils.RandomString(12),
		Logo: pgtype.Text{
			String: fmt.Sprintf("http://s3:url.%s.png", utils.RandomString(10)),
			Valid:  true,
		},
	}

	tenant, err := txTestStore.CreateTenant(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, tenant)
	require.NotEmpty(t, tenant.ID)
	require.Equal(t, tenant.Name, args.Name)
	require.Equal(t, tenant.Logo.String, args.Logo.String)

	return tenant
}

func TestCreateTenant(t *testing.T) {
	createRandomTenant(t)
}

func TestGetTenant(t *testing.T) {
	tenant := createRandomTenant(t)

	foundTenant, err := txTestStore.GetTenant(context.Background(), tenant.ID)

	require.NoError(t, err)
	require.Equal(t, tenant.ID, foundTenant.ID)
	require.Equal(t, tenant.Name, foundTenant.Name)
	require.Equal(t, tenant.Logo.String, foundTenant.Logo.String)
	require.WithinDuration(t, tenant.CreatedAt, foundTenant.CreatedAt, time.Second)
}
