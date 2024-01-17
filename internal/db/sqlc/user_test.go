package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:     util.RandomUserName(),
		Email:        util.RandomEmail(),
		FirstName:    util.RandomName(),
		PasswordHash: "testpassword",
		Role:         "test",
		HouseholdID:  pgtype.UUID{Bytes: [16]byte{}},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.HouseholdID, user.HouseholdID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.UserID)

	return user
}
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestFetchUserByUserName(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQueries.FetchUserByUserName(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEmpty(t, user2.Email)
	require.NotEmpty(t, user2.PasswordHash)
	require.NotEmpty(t, user2.Role)
	require.Equal(t, user1.Username, user2.Username)
	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)
	require.NotZero(t, user2.UserID)
}

func TestFetchUserByUserId(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQueries.FetchUserByUserId(context.Background(), user1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEmpty(t, user2.Email)
	require.NotEmpty(t, user2.PasswordHash)
	require.NotEmpty(t, user2.Role)
	require.NotEmpty(t, user2.Username)
	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)
	require.Equal(t, user1.UserID, user2.UserID)
}

func TestFetchUserByEmail(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQueries.FetchUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEmpty(t, user2.Username)
	require.NotEmpty(t, user2.PasswordHash)
	require.NotEmpty(t, user2.Role)
	require.Equal(t, user1.Email, user2.Email)
	require.NotZero(t, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)
	require.NotZero(t, user2.UserID)
}

func TestListHouseholdMembers(t *testing.T) {

}

func TestListUsers(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
