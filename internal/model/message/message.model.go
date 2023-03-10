package message

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID              *uuid.UUID     `json:"id" gorm:"primary_key"`
	Author          string         `json:"author"`
	Message         string         `json:"message"`
	Likes           uint           `json:"likes"`
	IsDeleted       gorm.DeletedAt `json:"isDeleted" gorm:"index;type:timestamp"`
	LastUpdateAt    *time.Time     `json:"lastUpdateAt" gorm:"type:timestamp;autoUpdateTime:nano"`
	LastImageUpdate *time.Time     `json:"lastImageUpdate" gorm:"type:timestamp"`
}

type CreateMessageDto struct {
	ID          string `json:"uuid" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Message     string `json:"message"`
	Likes       uint   `json:"likes"`
	ImageUpdate bool   `json:"imageUpdate"`
	Image       string `json:"image"`
}

type UpdateMessageDto struct {
	Author      string `json:"author" binding:"required"`
	Message     string `json:"message"`
	Likes       uint   `json:"likes"`
	ImageUpdate bool   `json:"imageUpdate"`
	Image       string `json:"image"`
}
