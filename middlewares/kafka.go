package middlewares

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-kafka-example/broker"
)

func KafkaMiddleware(producer sarama.SyncProducer, topic string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logMessage := broker.LogMessage{
			Method:     c.Request.Method,
			Endpoint:   c.Request.URL.Path,
			StatusCode: c.Writer.Status(),
			Timestamp:  time.Now().Format(time.RFC3339),
		}

		broker.SendLogToKafka(producer, topic, logMessage)
	}
}
