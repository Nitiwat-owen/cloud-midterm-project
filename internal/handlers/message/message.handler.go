package message

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/model/message"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func CreateMessage(c *gin.Context) {
	var requestBody message.MessageDto
	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(requestBody.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := &message.Message{
		ID:           &id,
		Author:       requestBody.Author,
		Message:      requestBody.Message,
		Likes:        requestBody.Likes,
		ImageUpdate:  requestBody.ImageUpdate,
		LastUpdateAt: time.Now(),
	}

	err = database.DB.Create(message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}
