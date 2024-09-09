package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	RPCPort  string
	KafkaUrl string

	Context struct {
		Timeout string
	}

	Mongo struct {
		Host     string
		Port     string
		User     string
		Password string
	}
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}

func New() *Config {
	var config Config

	// general configuration
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// Mongo configuration
	config.RPCPort = getEnv("RPC_PORT", ":50051")
	config.Mongo.Host = getEnv("MONGOHOST", "localhost")
	config.Mongo.Port = getEnv("MONGOPORT", "5432")
	config.Mongo.User = getEnv("MONGOUSER", "mubina")
	config.Mongo.Password = getEnv("MONGOPASSWORD", "1234")

	//db configuration
	config.DB.Host = getEnv("DBHOST", "localhost")
	config.DB.Port = cast.ToInt(getEnv("DBPORT", "5432"))
	config.DB.User = getEnv("DBUSER", "postgres")
	config.DB.Password = getEnv("DBPASSWORD", "1234")
	config.DB.Database = getEnv("DBDATABASE", "orders")

	config.KafkaUrl = getEnv("KAFKAURL", ":9095")
	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
