package config

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// struct config
type Config struct {
	Port     string         `env:"PORT" envDefault:"8080"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	Redis    RedisConfig    `envPrefix:"REDIS_"`
	JWT      JWTConfig      `envPrefix:"JWT_"`
}

// struct postgres
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}

// struct redis
type RedisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"6379"`
	Password string `env:"PASSWORD" envDefault:""`
}

// struct jwt
type JWTConfig struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"secret"`
}

// untuk membuat new config
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}

	// Tambahkan log ini untuk melihat apakah port terbaca dengan benar
	fmt.Printf("Server akan berjalan di port: %s\n", cfg.Port)

	if cfg.Port == "" {
		return nil, errors.New("port is not specified in the configuration")
	}

	return cfg, nil
}

// func untuk parse config
func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}
