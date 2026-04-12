package server

import (
	"log"

	"boilerplate/server/config"
	"boilerplate/server/services"
)

type Application struct {
	Config         config.Config
	Logger         *log.Logger
	ExampleService *services.ExampleService
}

func NewApplication(cfg config.Config) *Application {
	return &Application{
		Config:         cfg,
		Logger:         log.Default(),
		ExampleService: services.NewExampleService(),
	}
}
