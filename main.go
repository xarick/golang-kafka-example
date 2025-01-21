package main

import (
	"log"

	"github.com/xarick/golang-kafka-example/broker"
	"github.com/xarick/golang-kafka-example/config"
	"github.com/xarick/golang-kafka-example/db"
	"github.com/xarick/golang-kafka-example/middlewares"
	"github.com/xarick/golang-kafka-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectDB(cfg)
	defer db.DB.Close()

	kafkaProducer := broker.InitKafkaProducer(cfg.KafkaBroker)
	defer kafkaProducer.Close()

	r := routes.SetupRouter()
	r.Use(middlewares.KafkaMiddleware(kafkaProducer, cfg.KafkaTopic))

	go broker.StartConsumer(cfg)

	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
