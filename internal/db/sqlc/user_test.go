package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ShadrackAdwera/go-multitenancy/helpers"
	"github.com/ShadrackAdwera/go-utils/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	username := utils.RandomString(8)
	tenant := createRandomTenant(t)

	arg := CreateUserParams{
		Username: username,
		Email:    fmt.Sprintf("%s@mail.com", username),
		TenantID: tenant.ID,
		Password: hashedPassword,
	}

	user, err := txTestStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := txTestStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyUserName(t *testing.T) {
	oldUser := createRandomUser(t)

	newUserName := utils.RandomString(12)
	updatedUser, err := txTestStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUserName,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Password, updatedUser.Password)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := fmt.Sprintf("%s@mail.com", utils.RandomString(12))
	updatedUser, err := txTestStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.Password, updatedUser.Password)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := utils.RandomString(6)
	newHashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := txTestStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := createRandomUser(t)

	tenant := createRandomTenant(t)

	newUsername := utils.RandomString(12)
	newEmail := fmt.Sprintf("%s@mail.com", newUsername)
	newPassword := utils.RandomString(6)
	newHashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := txTestStore.UpdateUser(context.Background(), UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUsername,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
		TenantID: pgtype.UUID{
			Bytes: tenant.ID,
			Valid: true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, newHashedPassword, updatedUser.Password)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, newUsername, updatedUser.Username)
	require.NotEqual(t, oldUser.TenantID, updatedUser.TenantID)
	require.Equal(t, tenant.ID, updatedUser.TenantID)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := txTestStore.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)

	user1, err := txTestStore.GetUser(context.Background(), user.ID)

	require.Error(t, err)
	require.EqualError(t, err, helpers.ErrRecordNotFound.Error())
	require.Empty(t, user1)
}
