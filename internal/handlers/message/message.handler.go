package message

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/model/message"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateMessage(c *gin.Context) {
	var requestBody message.CreateMessageDto
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
		LastImageUpdate:  time.Now(),
		LastUpdateAt: time.Now(),
	}

	err = database.DB.Create(message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func UpdateMessage(c *gin.Context){
	var oldMessage message.Message
	if err := database.DB.Where("id = ?", c.Param("uuid")).First(&oldMessage).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
    return
  }

	var requestBody message.UpdateMessageDto
	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := uuid.Parse(c.Param("uuid"))
	
	update := &message.Message{
		ID:           &id,
		Author:       requestBody.Author,
		Message:      requestBody.Message,
		Likes:        requestBody.Likes,
	}
	if requestBody.ImageUpdate{
		update.LastImageUpdate = time.Now()
	}
	database.DB.Model(&oldMessage).Updates(update)
	c.JSON(204,gin.H{"data": oldMessage})
}

func DeleteMessage(c *gin.Context){
	var oldMessage message.Message
	if err := database.DB.Where("id = ?", c.Param("uuid")).First(&oldMessage).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
    return
  }
	database.DB.Delete(&oldMessage)
	c.JSON(204,gin.H{"data": true})
}