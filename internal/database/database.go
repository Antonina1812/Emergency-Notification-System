package database

import (
	"emergency_notification_system/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(red("Failed to connect to database:", err))
	}

	fmt.Println(green("Connected to database"))

	err = db.AutoMigrate(&models.User{}, &models.Device{}, &models.Message{}, &models.MessageHistory{}) // создание таблиц
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
}
