package config

import (
	"module/internal/models"
	"module/internal/services"
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

func TokenConfigLoadToService(cfg models.TokenConfig) {

	services.AccessKey = []byte(cfg.AccessKey)
	services.RefreshKey = []byte(cfg.RefreshKey)
	services.AccessDuration = cfg.AccessTTL
	services.RefreshDuration = cfg.RefreshTTL

	if cfg.AccessTTL == 0 || cfg.RefreshTTL == 0 {
		panic("empty TTL in token")
	}
	if cfg.AccessKey == "" || cfg.RefreshKey == "" {
		panic("empty keys in token")
	}

}
