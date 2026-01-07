package config

import "os"

type Env struct {
	Port   string
	DBName string
}

var Envs = initConfig()

func initConfig() *Env {
	return &Env{
		Port:   getEnv("PORT", "8080"),
		DBName: getEnv("DB_NAME", "db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
