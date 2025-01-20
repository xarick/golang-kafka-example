package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-kafka-example/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}

	return r
}
