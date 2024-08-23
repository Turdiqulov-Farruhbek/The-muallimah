package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_HTTP_PORT        string
	AUTH_GRPC_PORT        string
	DB_HOST               string
	DB_PORT               int
	DB_USER               string
	DB_PASSWORD           string
	DB_NAME               string
	SENDER_EMAIL          string
	APP_PASSWORD          string
	MONGO_URI             string
	MONGO_DB_NAME         string
	MONGO_COLLECTION_NAME string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_HTTP_PORT = cast.ToString(coalesce("AUTH_PORT", ":8088"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "root"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "delivery_auth"))
	config.SENDER_EMAIL = cast.ToString(coalesce("SENDER_EMAIL", "email@example.com"))
	config.APP_PASSWORD = cast.ToString(coalesce("APP_PASSWORD", "your_app_password_here"))
	config.MONGO_URI = cast.ToString(coalesce("MONGO_URI", "mongodb://mongo:root@mongo_db:27017"))
	config.MONGO_DB_NAME = cast.ToString(coalesce("MONGO_DB_NAME", "delivery_auth"))
	config.MONGO_COLLECTION_NAME = cast.ToString(coalesce("MONGO_COLLECTION_NAME", "users_data"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
