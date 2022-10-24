package sys

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env      string
	Port     string
	Postgres PostgresConfig
	Token    TokenConfig
}

type PostgresConfig struct {
	DSN string
}

type TokenConfig struct {
	AccessExpiration  time.Duration
	RefreshExpiration time.Duration
	PubKey            string
	PrivKey           string
}

func NewConfigWithEnv() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	accExp, err := strconv.Atoi(getEnv("ACCESS_EXP"))
	if err != nil {
		return nil, err
	}
	refExp, err := strconv.Atoi(getEnv("REFRESH_EXP"))
	if err != nil {
		return nil, err
	}

	conf := Config{
		Env:  getEnv("APP_MODE"),
		Port: getEnv("PORT"),
		Postgres: PostgresConfig{
			DSN: getEnv("POSTGRES_DSN"),
		},
		Token: TokenConfig{
			AccessExpiration:  time.Minute * time.Duration(accExp),
			RefreshExpiration: time.Hour * time.Duration(refExp),
			PubKey:            getEnv("PUB_KEY"),
			PrivKey:           getEnv("PRIV_KEY"),
		},
	}

	return &conf, nil
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("[KEY : %s] IS EMPTY", key))
	}
	return val
}
