package message

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID           *uuid.UUID     `json:"id" gorm:"primary_key"`
	Author       string         `json:"author"`
	Message      string         `json:"message"`
	Likes        uint           `json:"likes"`
	ImageUpdate  bool           `json:"imageUpdate"`
	IsDeleted    gorm.DeletedAt `json:"isDeleted" gorm:"index;type:timestamp"`
	LastUpdateAt time.Time      `json:"lastUpdateAt" gorm:"type:timestamp;autoUpdateTime:nano"`
}

type MessageDto struct {
	ID          string `json:"uuid" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Message     string `json:"message" binding:"required"`
	Likes       uint   `json:"likes"`
	ImageUpdate bool   `json:"imageUpdate"`
	Image       string `json:"image"`
}
