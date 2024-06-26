package configs

import (
	"go-marketplace/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:r!Z@L123@tcp(127.0.0.1:3306)/go_marketplace?parseTime=True"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database Connected")
}
