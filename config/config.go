package config

import (
	"gin-boilerplate/src/utils/logger"

	"github.com/caarlos0/env/v11"
)

type AppConfig struct {
	Port      string `env:"APP_PORT" envDefault:":3000"`
	Mode      string `env:"APP_MODE" envDefault:"prod"`
	Url       string `env:"APP_URL"`
}

type DBConfig struct {
	DSN      string `env:"DB_DSN"`
}

type Config struct {
	App AppConfig
	DB  DBConfig
}

func GetConfig() Config {

	cfg := Config{}
	opts := env.Options{RequiredIfNoDef: true}

	if err := env.ParseWithOptions(&cfg, opts); err != nil {
		logger.Fatal("Error during parse .env", "err", err.Error())
	}

	return cfg
}
