package config

import (
	"os"
	"strconv"
)

type Env struct {
	Port   string
	DBName string

	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() *Env {
	return &Env{
		Port:                   getEnv("PORT", "8080"),
		DBName:                 getEnv("DB_NAME", "db"),
		JWTExpirationInSeconds: getEnvInt("JWT_EXP", 10),
		JWTSecret:              getEnv("JWT_SECRET", "necesseraly_set_jwt_secret_in_env"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return v
	}

	return fallback
}
