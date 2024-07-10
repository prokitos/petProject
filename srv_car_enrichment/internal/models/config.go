package models

type Config struct {
	Env    string       `yaml:"env" env-default:"local"`
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
