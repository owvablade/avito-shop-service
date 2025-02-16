package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort       string `env:"APP_PORT" envDefault:"8080"`
	DbHost        string `env:"DB_HOST" envDefault:"localhost"`
	DbPort        int    `env:"DB_PORT"`
	DbUser        string `env:"DB_USER"`
	DbPassword    string `env:"DB_PASSWORD"`
	DbName        string `env:"DB_NAME" envDefault:"postgres"`
	DbSslMode     string `env:"DB_SSLMODE" envDefault:"disable"`
	DbDriver      string `env:"DB_DRIVER" envDefault:"postgres"`
	DbDialect     string `env:"DB_DIALECT" envDefault:"postgres"`
	DbMaxOpenConn int    `env:"DB_MAX_OPEN_CONN" envDefault:"5"`
	DbMaxIdleConn int    `env:"DB_MAX_IDLE_CONN" envDefault:"5"`
	MigrationsDir string `env:"MIGRATIONS_DIR" envDefault:"migrations"`
	JwtSecret     string `env:"JWT_SECRET"`
}

func LoadConfig(filename string) (*Config, error) {
	if err := godotenv.Load(filename); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
