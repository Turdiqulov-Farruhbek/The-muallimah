package config

import (
	"os"
)

type Config struct {
	RPCPort     string

	Context struct {
		Timeout string
	}

	DB struct {
		Host     string
		Port     string
		User     string
		Password string
	}
}

func New() *Config {
	var config Config

	// general configuration
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// db configuration
	config.RPCPort = getEnv("RPC_PORT", ":50051")
	config.DB.Host = getEnv("MONGOHOST", "localhost")
	config.DB.Port = getEnv("MONGOPORT", "5432")
	config.DB.User = getEnv("MONGOUSER", "mongo")
	config.DB.Password = getEnv("MONGOPASSWORD", "1234")
	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
