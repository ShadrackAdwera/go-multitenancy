package db

import (
	"context"
	"fmt"
	"testing"

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
