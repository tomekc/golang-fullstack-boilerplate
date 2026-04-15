package database

import (
	"io/fs"

	"github.com/jmoiron/sqlx"
	_ "github.com/ncruces/go-sqlite3/driver" // registers "sqlite3" driver
	"github.com/pressly/goose/v3"
)

func Init(path string, migrationsFS fs.FS) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := runMigrations(db, migrationsFS); err != nil {
		return nil, err
	}
	return db, nil
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
