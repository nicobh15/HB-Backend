package token

import (
	"time"
)

type Maker interface {
	CreateToken(user TokenableUser, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
