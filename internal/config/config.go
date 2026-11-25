package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN    string
	Secret []byte
	Addr   string
}

func DatabaseURL() (string, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return "", errors.New("DATABASE_URL wajib diisi")
	}
	return dsn, nil
}

func Load() (Config, error) {
	dsn, err := DatabaseURL()
	if err != nil {
		return Config{}, err
	}
	c := Config{DSN: dsn, Secret: []byte(os.Getenv("JWT_SECRET")), Addr: os.Getenv("HTTP_ADDR")}
	if len(c.Secret) < 32 {
		return c, errors.New("JWT_SECRET minimal 32 byte")
	}
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}
	return c, nil
}
