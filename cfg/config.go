package cfg

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server
}

type Server struct {
	Host        string `envconfig:"DML_HOST" default:"0.0.0.0"`
	Port        int    `envconfig:"DML_PORT" default:"8081"`
	EnableDebug bool   `envconfig:"DML_ENABLE_DEBUG" default:"true"`
}

func Parse() (Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
