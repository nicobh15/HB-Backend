package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type TokenableUser struct {
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	FirstName   string      `json:"first_name"`
	Role        string      `json:"role"`
	HouseholdID pgtype.UUID `json:"household_id"`
}
type Payload struct {
	ID        uuid.UUID     `json:"id"`
	User      TokenableUser `json:"user"`
	IssuedAt  time.Time     `json:"issued_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

func CastTokenableUser(user db.User) TokenableUser {
	return TokenableUser{
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		Role:        user.Role,
		HouseholdID: user.HouseholdID,
	}
}
func NewPayload(user TokenableUser, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		User:      user,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
