package utils

import (
	"log"
	"os"
)

type AppLogger struct {
	logger *log.Logger
}

var appLogger *AppLogger = nil

func NewLogger() *AppLogger {
	if appLogger != nil {
		return appLogger
	}
	logger := AppLogger{}
	logger.logger = log.New(os.Stdout, "library-app-api", log.LstdFlags|log.Lshortfile)
	appLogger = &logger
	return appLogger
}

func (logger *AppLogger) Info(message string) {
	logger.logger.Println(message)
}

func (logger *AppLogger) Error(message string) {
	logger.logger.Println(message)
}

func (logger *AppLogger) Fatal(message string) {
	logger.logger.Fatalln(message)
}

func (logger *AppLogger) Panic(message string) {
	logger.logger.Panicln(message)
}
