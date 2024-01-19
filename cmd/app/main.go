package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicobh15/HomeBuddy-Backend/internal/api"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/homebuddy?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	pool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(pool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
