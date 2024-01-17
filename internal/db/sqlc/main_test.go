package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

const dbSource = "postgresql://root:secret@localhost:5432/homebuddy?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())

}
