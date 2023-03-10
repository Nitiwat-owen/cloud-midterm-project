package message

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/model/message"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
		ID:      &id,
		Author:  requestBody.Author,
		Message: requestBody.Message,
		Likes:   requestBody.Likes,
	}
	if requestBody.ImageUpdate {
		currentTime, _ := time.Parse(time.Layout, time.Now().Format(time.Layout))
		message.LastImageUpdate = &currentTime
	}

	err = database.DB.Create(message).Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if requestBody.ImageUpdate {
		filename := fmt.Sprintf("%s.txt", requestBody.ID)
		fmt.Println(filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(requestBody.Image)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": message})
}

func UpdateMessage(c *gin.Context) {
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
		ID:      &id,
		Author:  requestBody.Author,
		Message: requestBody.Message,
		Likes:   requestBody.Likes,
	}
	if requestBody.ImageUpdate {
		currentTime, _ := time.Parse(time.Layout, time.Now().Format(time.Layout))
		update.LastImageUpdate = &currentTime
	}
	database.DB.Model(&oldMessage).Updates(update)
	c.JSON(204, gin.H{"data": oldMessage})
}

func DeleteMessage(c *gin.Context) {
	var oldMessage message.Message
	if err := database.DB.Where("id = ?", c.Param("uuid")).First(&oldMessage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&oldMessage)
	c.JSON(204, gin.H{"data": true})
}
