package logger

import (
	"github.com/google/uuid"
	"log"
)

// Logger contains all log related data
// used when logging reuqest messages
type Logger struct {
	ID string
}

//NewLogger creates a new logger
func NewLogger() *Logger {
	id := uuid.New()
	return &Logger{ID: id.String()}
}

// Log logs message in the following format: "type [ID]: contents"
func (l *Logger) Log(logType string, message string) {
	log.Printf("%s [%s]: %s", logType, l.ID, message)
}
