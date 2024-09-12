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
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
}

func New() *Config {
	var config Config

	// General configuration
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// MongoDB configuration
	config.RPCPort = getEnv("RPC_PORT", ":50051")
	config.Mongo.Host = getEnv("MONGOHOST", "mongodb")
	config.Mongo.Port = getEnv("MONGOPORT", "27017")
	config.Mongo.User = getEnv("MONGOUSER", "root")
	config.Mongo.Password = getEnv("MONGOPASSWORD", "root")

	// PostgreSQL configuration
	config.PostgresHost = cast.ToString(getEnv("POSTGRES_HOST", "muallimah_db"))
	config.PostgresPort = cast.ToInt(getEnv("POSTGRES_PORT", "5432"))
	config.PostgresUser = cast.ToString(getEnv("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getEnv("POSTGRES_PASSWORD", "00salom00"))
	config.PostgresDatabase = cast.ToString(getEnv("POSTGRES_DATABASE", "muallimah_db"))

	// Kafka configuration
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
