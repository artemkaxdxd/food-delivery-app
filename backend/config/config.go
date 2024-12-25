package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Postgres Postgres
		Server   Server
		JWT      JWT
		IsDev    bool `envconfig:"IS_DEV"`
	}

	Postgres struct {
		Host     string `envconfig:"POSTGRES_HOST"`
		Port     string `envconfig:"POSTGRES_PORT"`
		User     string `envconfig:"POSTGRES_USER"`
		Password string `envconfig:"POSTGRES_PASSWORD"`
		DB       string `envconfig:"POSTGRES_DB"`
	}

	Server struct {
		Port string `envconfig:"SERVER_PORT"`
	}

	JWT struct {
		Secret string `envconfig:"JWT_SECRET"`
	}
)

func FillConfig() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
