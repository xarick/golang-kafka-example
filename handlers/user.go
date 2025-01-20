package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-kafka-example/db"
)

func Users(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
