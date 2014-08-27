package bunyan

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	Fatal LogLevel = 60
	Error          = 50
	Warn           = 40
	Info           = 30
	Debug          = 20
	Trace          = 10
)

type LogEntry struct {
	Data      interface{} `json:"data,omitempty"`
	Hostname  string      `json:"hostname"`
	Level     LogLevel    `json:"level"`
	Message   string      `json:"message"`
	Name      string      `json:"name"`
	ProcessID int         `json:"pid"`
	Time      time.Time   `json:"time"`
	Version   int         `json:"v"`
}

func hostname() string {
	result, err := os.Hostname()

	if err != nil {
		panic(fmt.Sprintf("Error retrieving hostname %v", err))
	}

	return result
}

var logEntryTemplate = LogEntry{
	Version:   0,
	Hostname:  hostname(),
	ProcessID: os.Getpid(),
}

func NewLogEntry(level LogLevel, message string, data interface{}) *LogEntry {
	result := logEntryTemplate

	result.Data = data
	result.Level = level
	result.Message = message
	result.Time = time.Now()

	return &result
}

func (l *LogEntry) setLogger(logger *Logger) {
	l.Name = logger.name
}

func (l *LogEntry) String() string {
	result, _ := json.Marshal(l)

	return string(result)
}
