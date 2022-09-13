package config

import (
	"example.choreo.dev/internal/models"
)

type Config struct {
	Env             string
	Port            int
	MaxPetsAllowed  int
	InitialDataPath string
}

type InitialData struct {
	Pets       []models.Pet      `json:"pets"`
	Categories []models.Category `json:"categories"`
}
