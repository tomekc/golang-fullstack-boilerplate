package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml/v2"

	"boilerplate/server/database"
)

// Parent folder of data directory: for configuration, caches, etc
const dataDirectory = "data"

type Config struct {
	DataDir  string
	Server   Server          `toml:"server"`
	Database database.Config `toml:"database"`
}

type Server struct {
	Port int `toml:"port"`
}

func Load() Config {
	dataDir := flag.String("datadir", "./"+dataDirectory, "Path to the data directory")
	flag.Parse()

	cfg := defaultConfig(*dataDir)
	// parse config file
	configFilePath := fmt.Sprintf("%s/config.toml", *dataDir)
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = toml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	log.Printf("Configuration loaded from: %s", configFilePath)

	return cfg
}

func defaultConfig(dataDir string) Config {
	return Config{
		DataDir:  dataDir,
		Database: database.Config{DSN: dataDir + "/db.sqlite3"},
	}
}
