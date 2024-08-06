package models

type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	Server   ServerConfig
	External ExternalConfig
}

type ExternalConfig struct {
	AuthService    string `yaml:"auth"`
	CarService     string `yaml:"car"`
	RabbitMqServer string `yaml:"rabbit"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}
