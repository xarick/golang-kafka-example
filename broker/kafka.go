package broker

import (
	"log"

	"github.com/IBM/sarama"
)

func InitKafkaProducer(broker string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}

	return producer
}
