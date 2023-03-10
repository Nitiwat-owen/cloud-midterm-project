package user

import "time"

type User struct {
	Username     string    `json:"username" gorm:"primary_key"`
	LastOnlineAt time.Time `json:"lastOnlineAt`
}
