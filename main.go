package main

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/handlers/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	database.ConnectDatabase()

	r.POST("/api/messages", message.CreateMessage)

	r.Run()
}
