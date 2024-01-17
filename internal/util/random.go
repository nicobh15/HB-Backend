package util

import (
	"math/rand"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func RandomInt(min, max int64) int64 {
	return (min + rand.Int63n(max-min+1))
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return (RandomString(10) + "@test.com")
}

func RandomUserName() string {
	return (RandomString(10))
}

func RandomName() string {
	return (RandomString(7))
}

func RandomAddress() pgtype.Text {
	return (pgtype.Text{String: RandomString(10), Valid: true})
}
