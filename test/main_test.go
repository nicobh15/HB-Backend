package test

import (
	"context"
	"log"
	"os"
	"testing"

	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"

	"github.com/jackc/pgx/v5"
)

var testQueries *db.Queries

const dbSource = "postgresql://root:secret@localhost:5432/homebuddy?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(conn)
	os.Exit(m.Run())

}
