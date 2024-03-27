package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nicobh15/hb-backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(10))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomUserName(),
		Email:        util.RandomEmail(),
		FirstName:    util.RandomName(),
		PasswordHash: hashedPassword,
		Role:         util.RandomName(),
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
	household := CreateRandomHousehold(t)

	for i := 0; i < 5; i++ {
		user := CreateRandomUser(t)
		arg := UpdateUserParams{
			Username:     user.Username,
			Email:        user.Email,
			FirstName:    user.FirstName,
			PasswordHash: user.PasswordHash,
			Role:         user.Role,
			HouseholdID:  household.HouseholdID,
			UserID:       user.UserID}
		testQueries.UpdateUser(context.Background(), arg)
	}
	args := ListHouseholdMembersParams{
		HouseholdID: household.HouseholdID,
		Limit:       5,
		Offset:      0,
	}

	users, err := testQueries.ListHouseholdMembers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, users, 5)
	require.NotEmpty(t, users)

}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	household := CreateRandomHousehold(t)

	arg := UpdateUserParams{
		Username:     util.RandomName(),
		Email:        util.RandomEmail(),
		FirstName:    util.RandomName(),
		PasswordHash: util.RandomName(),
		Role:         util.RandomName(),
		HouseholdID:  household.HouseholdID,
		UserID:       user1.UserID}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotEqual(t, user2.Username, user1.Username)
	require.NotEqual(t, user2.Email, user1.Email)
	require.NotEqual(t, user2.FirstName, user1.FirstName)
	require.NotEqual(t, user2.PasswordHash, user1.PasswordHash)
	require.NotEqual(t, user2.HouseholdID, user1.HouseholdID)
	require.NotEqual(t, user2.Role, user1.Role)
	require.NotZero(t, user2.CreatedAt)
	require.WithinDuration(t, user2.UpdatedAt.Time, user1.UpdatedAt.Time, time.Second)

}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	_, err := testQueries.DeleteUser(context.Background(), user1.Email)

	user2, err2 := testQueries.FetchUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.Error(t, err2)
	require.EqualError(t, err2, "no rows in result set")
	require.Empty(t, user2)
}
