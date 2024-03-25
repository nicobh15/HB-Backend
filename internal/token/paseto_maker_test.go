package token

import (
	"testing"
	"time"

	"github.com/nicobh15/hb-backend/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateUser() TokenableUser {

	user := TokenableUser{
		Username:    util.RandomUserName(),
		Email:       util.RandomEmail(),
		FirstName:   util.RandomName(),
		Role:        util.RandomName(),
		HouseholdID: util.RandomUUID(),
	}

	return user
}
func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	user := CreateUser()

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(user, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.User.Username, user.Username)
	require.Equal(t, payload.User.Email, user.Email)
	require.Equal(t, payload.User.FirstName, user.FirstName)
	require.Equal(t, payload.User.Role, user.Role)
	require.Equal(t, payload.User.HouseholdID, user.HouseholdID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	user := CreateUser()
	duration := time.Minute

	token, err := maker.CreateToken(user, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}
