package message

import (
	"cloud-midterm-project/database"
	"cloud-midterm-project/internal/model/message"
	user "cloud-midterm-project/internal/model/user"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMessage(c *gin.Context) {
	username := c.GetHeader("username")
	// query lastOnlineAt
	user := &user.User{}
	err := database.DB.Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// create a new user
			user.Username = username
			database.DB.Create(user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	messages := &[]message.Message{}
	err = database.DB.Where("last_update_at > ?", user.LastOnlineAt).Find(messages).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// update lastOnlineAt
	user.LastOnlineAt = time.Now()
	_ = database.DB.Model(user).Updates(user).Error

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

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
	id := c.Param("uuid")
	if err := database.DB.Where("id = ?", id).First(&oldMessage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	var requestBody message.UpdateMessageDto
	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := &message.Message{
		Author:  requestBody.Author,
		Message: requestBody.Message,
		Likes:   requestBody.Likes,
	}

	if requestBody.ImageUpdate {
		currentTime, _ := time.Parse(time.Layout, time.Now().Format(time.Layout))
		update.LastImageUpdate = &currentTime
	}
	database.DB.Model(&oldMessage).Updates(update)

	if requestBody.ImageUpdate {
		filename := fmt.Sprintf("%s.txt", id)

		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		if requestBody.Image == "" {
			err := os.Remove(filename)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(204, gin.H{"data": oldMessage})
			return
		}
		_, err = file.WriteString(requestBody.Image)
	}
	c.JSON(204, gin.H{"data": oldMessage})
}

func DeleteMessage(c *gin.Context) {
	var oldMessage message.Message
	id := c.Param("uuid")
	if err := database.DB.Where("id = ?", id).First(&oldMessage).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	database.DB.Delete(&oldMessage)
	filename := fmt.Sprintf("%s.txt", id)
	_, err := os.Stat(filename)
	if os.IsExist(err) {
		err := os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(204, gin.H{"data": true})
}
