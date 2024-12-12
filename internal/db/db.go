package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *pgxpool.Pool {
	dbURL := "postgres://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_URL") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_DATABASE") + "?sslmode=disable"

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Println("Error with DB initialization")
		os.Exit(1)
	}

	return dbPool
}
