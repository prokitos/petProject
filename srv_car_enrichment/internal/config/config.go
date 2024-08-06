package config

import (
	"module/internal/models"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func ConfigMustLoad(name string) *models.Config {

	path := "internal/config/" + name + ".yaml"
	var cfg models.Config

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("nothing from this path")
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read")
	}

	return &cfg
}
