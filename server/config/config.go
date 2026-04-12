package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

// Parent folder of data directory: for configuration, caches, etc
const dataDirectory = "data"

type Config struct {
	DataDir string
	Server  Server `toml:"server"`
}

type Server struct {
	Port int `toml:"port"`
}

func isTruthy(s string) bool {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "yes", "YES", "Yes", "on", "ON", "On":
		return true
	}
	return false
}

func Load() Config {
	devMode := flag.Bool("dev", false, "Development mode: runs server only on port ")
	dockerMode := flag.Bool("docker", false, "Docker mode: self contained, run in container")
	flag.Parse()

	dataDir := determineRunMode(*devMode || isTruthy(os.Getenv("DEV")), *dockerMode)
	cfg := defaultConfig(dataDir)
	// parse config file
	configFilePath := fmt.Sprintf("%s/config.toml", dataDir)
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
		DataDir: dataDir,
	}
}

func determineRunMode(devMode bool, dockerMode bool) string {
	if devMode && dockerMode {
		log.Fatalf("Illegal combination of dev and docker mode")
	}
	if devMode {
		return "./" + dataDirectory
	} else if dockerMode {
		log.Printf("Running in docker mode")
		return "/" + dataDirectory
	} else {
		log.Printf("Running in standalone mode")
		return "./" + dataDirectory
	}
}
