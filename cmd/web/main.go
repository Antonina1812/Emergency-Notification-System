package main

import (
	"emergency_notification_system/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var green = color.New(color.FgGreen).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(red("Error loading .env file"))
	}

	database.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Emergency Notification System"))
	})

	fmt.Println(green("Server starting on port", port))
	http.ListenAndServe(":"+port, nil)

}
