package bunyan

import (
	"fmt"
)

type Logger struct {
	name    string
	streams []StreamInterface
}

// Creates a new logger, given one or more streams
func NewLogger(name string, streams []StreamInterface) *Logger {
	return &Logger{
		name:    name,
		streams: streams,
	}
}

func (l *Logger) AddStream(s StreamInterface) {
	l.streams = append(l.streams, s)
}

func (l *Logger) Log(e *LogEntry) {
	e.setLogger(l)

	for _, stream := range l.streams {
		go stream.Publish(e)
	}
}

func (l *Logger) Logln(level LogLevel, message string) *LogEntry {
	e := NewLogEntry(level, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) LogF(level LogLevel, format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(level, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) LogD(level LogLevel, message string, data interface{}) *LogEntry {
	e := NewLogEntry(level, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Trace(message string) *LogEntry {
	e := NewLogEntry(Trace, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) TraceF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Trace, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) TraceD(message string, data interface{}) *LogEntry {

	e := NewLogEntry(Trace, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Debug(message string) *LogEntry {
	e := NewLogEntry(Debug, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) DebugF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Debug, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) DebugD(message string, data interface{}) *LogEntry {
	e := NewLogEntry(Debug, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Info(message string) *LogEntry {
	e := NewLogEntry(Info, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) InfoF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Info, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) InfoD(message string, data interface{}) *LogEntry {
	e := NewLogEntry(Info, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Warn(message string) *LogEntry {
	e := NewLogEntry(Warn, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) WarnF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Warn, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) WarnD(message string, data interface{}) *LogEntry {
	e := NewLogEntry(Warn, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Error(message string) *LogEntry {
	e := NewLogEntry(Error, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) ErrorF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Error, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) ErrorD(message string, data interface{}) *LogEntry {
	e := NewLogEntry(Error, message, data)

	l.Log(e)

	return e
}

func (l *Logger) Fatal(message string) *LogEntry {
	e := NewLogEntry(Fatal, message, nil)

	l.Log(e)

	return e
}

func (l *Logger) FatalF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Fatal, fmt.Sprintf(format, values...), nil)

	l.Log(e)

	return e
}

func (l *Logger) FatalD(message string, data interface{}) *LogEntry {
	e := NewLogEntry(Fatal, message, data)

	l.Log(e)

	return e
}
