package config

import (
	"log"
	"os"

	"github.com/spf13/cast"

	"github.com/joho/godotenv"
)

type Config struct {
	GIN_PORT string

	DB_PORT     string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	CASBIN_DB_PORT string
	CASBIN_DB_HOST string
	CASBIN_DB_NAME string

	SIGNING_KEY  string
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{}

	config.GIN_PORT = cast.ToString(coalesce("GIN_PORT", ":8080"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "postgres"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "road_24"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "BEKJONS"))

	config.CASBIN_DB_HOST = cast.ToString(coalesce("CASBIN_DB_HOST", "localhost"))
	config.CASBIN_DB_PORT = cast.ToString(coalesce("CASBIN_DB_PORT", "5432"))
	config.CASBIN_DB_NAME = cast.ToString(coalesce("CASBIN_DB_NAME", "casbin"))

	config.SIGNING_KEY = cast.ToString(coalesce("SIGNING_KEY", "SSECCA"))

	return config

}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
