package logger

import (
	"log"
)

// LogError logs a fatal error
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
