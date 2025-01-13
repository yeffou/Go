package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func Init() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logger = log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	logger.Println("INFO: " + message)
}

func LogError(err error) {
	if err != nil {
		logger.Println("ERROR: " + err.Error())
	}
}
