package services

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// TODO: need to check state of migrations
func UpDatabase(connectionStr string) error {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to getting migrations: %w", err)
	}

	_ = m.Up()
	return nil
}
