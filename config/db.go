package config

import (
	"belajar-gin/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB
func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/cobaa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB COnnection Error:", err)
	}

	DB = db

	db.AutoMigrate(&models.User{})
	log.Printf("Connect")
}
