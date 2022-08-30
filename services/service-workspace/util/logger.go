package util

import (
	"log"
	"os"
)

type LogLevel int16

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

type Logger struct {
	Log *log.Logger
}

func (logger *Logger) Debug(message string, args ...interface{}) {
	logger.Log.Printf(message, args...)
}

func (logger *Logger) Info(message string, args ...interface{}) {
	logger.Log.Printf(message, args...)
}

func (logger *Logger) Warn(message string, args ...interface{}) {
	logger.Log.Printf(message, args...)
}

func (logger *Logger) Error(message error) {
	logger.Log.Print(message)
}

func (logger *Logger) ErrorS(message string, args ...interface{}) {
	logger.Log.Print(message)
}

func (logger *Logger) Fatal(message error) {
	logger.Log.Fatal(message)
}

func New() *Logger {
	return &Logger{
		Log: log.New(os.Stdout, "[service-workspace-api]", log.LstdFlags),
	}
}
