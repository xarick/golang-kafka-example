package main

import (
	"log"

	"github.com/xarick/golang-kafka-example/config"
	"github.com/xarick/golang-kafka-example/db"
	"github.com/xarick/golang-kafka-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg)
	defer db.DB.Close()

	r := routes.SetupRouter()
	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
