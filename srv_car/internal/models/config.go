package models

type Config struct {
	Env     string        `yaml:"env" env-default:"local"`
	Connect ConnectConfig `yaml:"postgres"`
	Server  ServerConfig  `yaml:"server"`
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
