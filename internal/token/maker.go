package token

import (
	"time"

	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

type Maker interface {
	CreateToken(user db.User, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
