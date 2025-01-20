package broker

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/xarick/golang-kafka-example/config"
)

func StartConsumer(cfg config.Application) {
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaBroker}, nil)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Println("Consumer is listening for messages...")
	for message := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
