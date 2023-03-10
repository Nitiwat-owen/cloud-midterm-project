package main

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/handlers/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	database.ConnectDatabase()

	r.POST("/api/messages", message.CreateMessage)
	r.PUT("/api/messages/:uuid", message.UpdateMessage)
	r.DELETE("/api/messages/:uuid", message.DeleteMessage)

	r.Run()
}
