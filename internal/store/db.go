package store

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DBConfig holds the configuration for the database connection
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewDB initializes a new database connection using pgx
func NewDB(cfg DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func main() {
	cfg := DBConfig{
		Host:     "test-db-1.cm5eoyii3fwi.us-east-1.rds.amazonaws.com",
		Port:     5432,
		User:     "postgres",
		Password: "vlnwXIHjwqjqWjZcCxJU",
		SSLMode:  "require",
	}

	dbpool, err := NewDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbpool.Close()

	// Your database operations go here
}
