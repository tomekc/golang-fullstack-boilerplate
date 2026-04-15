package server

import (
	"log"

	"boilerplate/migrations"
	"boilerplate/server/config"
	"boilerplate/server/database"
	"boilerplate/server/services"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	Config         config.Config
	Logger         *log.Logger
	DB             *sqlx.DB
	ExampleService *services.ExampleService
}

func NewApplication(cfg config.Config) *Application {
	db, err := database.Init(cfg.Database.Path, migrations.FS)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	return &Application{
		Config:         cfg,
		Logger:         log.Default(),
		DB:             db,
		ExampleService: services.NewExampleService(),
	}
}
