package routes

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-kafka-example/handlers"
	"github.com/xarick/golang-kafka-example/middlewares"
)

func SetupRouter(producer sarama.SyncProducer, topic string) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.KafkaMiddleware(producer, topic))

	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}

	r.GET("/users", handlers.Users)

	return r
}
