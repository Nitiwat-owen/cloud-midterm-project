package database

import (
	"cloud-midterm-project/config"
	"cloud-midterm-project/internal/model/message"
	"cloud-midterm-project/internal/model/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("fail to connect database")
	}

	DB.AutoMigrate(&message.Message{}, &user.User{})
	fmt.Println("Connection Opened to Database")

}
