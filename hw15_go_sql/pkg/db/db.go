package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connect(ctx context.Context, _ string) (*pgxpool.Pool, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var db *pgxpool.Pool
	db, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err = db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping databse: %w", err)
	}
	return db, nil
}
