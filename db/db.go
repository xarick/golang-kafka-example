package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xarick/golang-kafka-example/config"
)

var DB *sqlx.DB

func ConnectDB(cfg config.Application) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tashkent", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDbName)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	log.Println("Connected to PostgreSQL")
}
