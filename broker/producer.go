package broker

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type LogMessage struct {
	Method     string `json:"method"`
	Endpoint   string `json:"endpoint"`
	StatusCode int    `json:"status_code"`
	Timestamp  string `json:"timestamp"`
}

func SendLogToKafka(producer sarama.SyncProducer, topic string, logMessage LogMessage) {
	messageBytes, err := json.Marshal(logMessage)
	if err != nil {
		log.Printf("Failed to marshal log message: %v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send log message to Kafka: %v", err)
	}
}
