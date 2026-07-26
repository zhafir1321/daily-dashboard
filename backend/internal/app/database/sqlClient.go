package database

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

type Config struct {
	DBDriver          string
	DBSource          string
	MaxOpenConns      int
	MaxIdleConns      int
	ConnMaxIdleTime   time.Duration
	ConnectionTimeout time.Duration
}

type SQLClient struct {
	DB *sql.DB
}

//go:embed migrations/*.sql
var migrationsFS embed.FS

// NewSQLClient creates a new database client with the given configuration.
func NewSQLClient(cfg Config) (*SQLClient, error) {
	db, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	// Ping the database to verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnectionTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return &SQLClient{DB: db}, nil
}

func RunMigrations(databaseURL string) error {
	sourceDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create source driver: %w", err)
	}

	// Create a new migrate instance
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	return fmt.Errorf("failed to create database driver: %w", err)
	// }
	m, err := migrate.NewWithSourceInstance("iofs", sourceDriver, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	// Run the migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, _, _ := m.Version()
	log.Printf("Migrations completed. Current version: %d", version)

	return nil
}

// Close terminates the database connection.
func (client *SQLClient) Close() error {
	if client.DB != nil {
		return client.DB.Close()
	}
	return nil
}
