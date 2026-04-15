package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/ncruces/go-sqlite3/driver" // registers "sqlite3" driver
)

func Init(path string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
