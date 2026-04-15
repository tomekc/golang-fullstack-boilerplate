package database

import (
	"context"
	"io/fs"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/ncruces/go-sqlite3/driver" // registers "sqlite3" driver
	"github.com/pressly/goose/v3"
)

type Config struct {
	DSN string `toml:"dsn"`
}

type Database struct {
	*sqlx.DB
}

func New(cfg Config, migrationsFS fs.FS) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "sqlite3", cfg.DSN)
	if err != nil {
		return nil, err
	}
	if err := runMigrations(db, migrationsFS); err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

func runMigrations(db *sqlx.DB, migrationsFS fs.FS) error {
	goose.SetBaseFS(migrationsFS)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	if err := goose.Up(db.DB, "."); err != nil {
		return err
	}
	return nil
}
