package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

const (
	DefaultPort = 8080
)

var (
	EnvName         = "ENV"
	initialDataPath = "INIT_DATA_PATH"
	MaxPetsAllowed  = "MAX_PETS_ALLOWED"
)

var config Config

func GetConfig() *Config {
	return &config
}

func LoadConfig() (*Config, error) {
	config = Config{
		Port:            DefaultPort,
		Env:             os.Getenv(EnvName),
		InitialDataPath: os.Getenv(initialDataPath),
		MaxPetsAllowed:  getEnvInt(MaxPetsAllowed, 5),
	}
	return &config, nil
}

func LoadInitialData() (data InitialData) {
	if config.InitialDataPath == "" {
		return
	}
	contents, err := os.ReadFile(config.InitialDataPath)
	if err != nil {
		log.Fatalf("failed to read initial data at [%s]: %s", config.InitialDataPath, err)
	}
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatalf("failed to unmarshal initial data at [%s]: %s", config.InitialDataPath, err)
	}
	return
}

func getEnvInt(key string, defaultVal int) int {
	s := os.Getenv(key)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return v
}
