package test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
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
}

func TestFetchUserByUserName(t *testing.T) {
	Username := "testusername"

	user, err := testQueries.FetchUserByUserName(context.Background(), Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.PasswordHash)
	require.NotEmpty(t, user.Role)
	require.Equal(t, Username, user.Username)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.UserID)
}

func TestFetchUserByUserId(t *testing.T) {
	parsedUUID, err := uuid.Parse("10d36113-c6a9-4708-b84e-a0d076703359")
	if err != nil {
		t.Fatalf("Failed to parse UUID: %v", err)
	}

	UserID := pgtype.UUID{Bytes: parsedUUID, Valid: true}
	user, err := testQueries.FetchUserByUserId(context.Background(), UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Email)
	require.NotEmpty(t, user.PasswordHash)
	require.NotEmpty(t, user.Role)
	require.NotEmpty(t, user.Username)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.Equal(t, UserID, user.UserID)
}

func TestFetchUserByEmail(t *testing.T) {
	Email := "testuser@test.com"

	user, err := testQueries.FetchUserByEmail(context.Background(), Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotEmpty(t, user.Username)
	require.NotEmpty(t, user.PasswordHash)
	require.NotEmpty(t, user.Role)
	require.Equal(t, Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.UserID)
}

func TestListHouseholdMembers(t *testing.T) {

}

func TestListUsers(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
