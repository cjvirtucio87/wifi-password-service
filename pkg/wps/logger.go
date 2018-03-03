package wps

import (
	"log"
	"time"
)

type WpsLogger struct {
	logLevel string
}

func timestamp() string {
	return time.Now().Format(time.Now().Format("20060102150405"))
}

func NewWpsLogger(logLevel string) WpsLogger {
	return WpsLogger{
		logLevel: logLevel,
	}
}

func (logger WpsLogger) Info(m string) {
	log.Printf("[INFO] (%s)  %s\n", timestamp(), m)
}

func (logger WpsLogger) Debug(m string) {
	if logger.logLevel == "debug" {
		log.Printf("[DEBUG] (%s)  %s\n", timestamp(), m)
	}
}

func (logger WpsLogger) Error(m string) {
	log.Printf("[ERROR] (%s) %s\n", timestamp(), m)
}
