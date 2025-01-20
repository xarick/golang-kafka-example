package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode string
	RunPort string

	PgHost   string
	PgPort   string
	PgUser   string
	PgPass   string
	PgDbName string

	KafkaBroker string
	KafkaTopic  string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")

	cfg.PgHost = os.Getenv("PG_HOST")
	cfg.PgPort = os.Getenv("PG_PORT")
	cfg.PgUser = os.Getenv("PG_USER")
	cfg.PgPass = os.Getenv("PG_PASS")
	cfg.PgDbName = os.Getenv("PG_DB_NAME")

	cfg.KafkaBroker = os.Getenv("KAFKA_BROKER")
	cfg.KafkaTopic = os.Getenv("KAFKA_TOPIC")

	return cfg
}
