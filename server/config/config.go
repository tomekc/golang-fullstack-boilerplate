package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type RunMode int

// Parent folder of data directory: for configuration, caches, etc
const dataDirectory = "data"

const (
	RunModeStandalone = iota
	RunModeDevelopment
	RunModeDocker
)

type Config struct {
	RunMode RunMode
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

	dataDir, mode := determineRunMode(*devMode || isTruthy(os.Getenv("DEV")), *dockerMode)
	cfg := defaultConfig(dataDir, mode)
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

func defaultConfig(dataDir string, mode RunMode) Config {
	return Config{
		RunMode: mode,
		DataDir: dataDir,
	}
}

func determineRunMode(devMode bool, dockerMode bool) (string, RunMode) {
	if devMode && dockerMode {
		log.Fatalf("Illegal combination of dev and docker mode")
	}
	if devMode {
		return "./" + dataDirectory, RunModeDevelopment
	} else if dockerMode {
		log.Printf("Running in docker mode")
		return "/" + dataDirectory, RunModeDocker
	} else {
		log.Printf("Running in standalone mode")
		return "./" + dataDirectory, RunModeStandalone
	}
}
