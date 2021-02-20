package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// ConnectDatabase ...
func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", os.Getenv("DB_UNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection to database failed")
	}

	// Auto Migration
	database.AutoMigrate(&Todo{})

	DB = database
}
