package server

import (
	"log"

	"boilerplate/migrations"
	"boilerplate/server/config"
	"boilerplate/server/database"
	"boilerplate/server/services"
)

type Application struct {
	Config         config.Config
	Logger         *log.Logger
	DB             *database.Database
	ExampleService *services.ExampleService
}

func NewApplication(cfg config.Config) *Application {
	db, err := database.New(cfg.Database, migrations.FS)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	return &Application{
		Config:         cfg,
		Logger:         log.Default(),
		DB:             db,
		ExampleService: services.NewExampleService(db.DB),
	}
}
