package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(connStr string) *Repository {
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	return &Repository{
		DB: db,
	}
}
