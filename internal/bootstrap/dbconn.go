// Package bootstrap stores basic common entities such as config and database connection
package bootstrap

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

const (
	maxOpenConns      = 10
	reconnectAttempts = 3
	reconnectInterval = 3 * time.Second
)

// NewDBConn creates a new database connection (connection pool) instance.
func NewDBConn(username, password, name, host, port string) (*dbr.Connection, error) {
	// opening new connection; it's NOT necessary to close it
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username,
		password,
		name,
		host,
		port,
	)

	p, err := dbr.Open("postgres", dsn, nil)
	if err != nil {
		return nil, err
	}

	for c := reconnectAttempts; c > 0; c-- {
		err = p.Ping()
		if err != nil {
			time.Sleep(reconnectInterval)

			continue
		}

		break
	}

	if err != nil {
		return nil, err
	}

	// a maximum of 10 concurrent connections; might be changed later if needed
	p.SetMaxOpenConns(maxOpenConns)

	return p, nil
}

func UpMigrations(db *sql.DB, dbName, path string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create driver instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		dbName, driver)
	if err != nil {
		return fmt.Errorf("create a new migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("apply migrations: %w", err)
	}

	return nil
}
