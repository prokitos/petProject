package models

import "time"

type Config struct {
	Env     string        `yaml:"env" env-default:"local"`
	Connect ConnectConfig `yaml:"postgres"`
	Server  ServerConfig  `yaml:"server"`
	Token   TokenConfig   `yaml:"token"`
}

type ConnectConfig struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Name   string `yaml:"name"`
	Reload bool   `yaml:"reload"`
	Port   string `yaml:"port"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type TokenConfig struct {
	AccessTTL  time.Duration `yaml:"access_ttl"`
	RefreshTTL time.Duration `yaml:"refresh_ttl"`
	AccessKey  string        `yaml:"access_key"`
	RefreshKey string        `yaml:"refresh_key"`
}
