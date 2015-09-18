package bunyan

import (
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
	name        string
	streams     []StreamInterface
	sync        bool
	minLogLevel LogLevel
}

// Creates a new logger, given one or more streams
func NewLogger(name string, streams []StreamInterface) *Logger {
	return &Logger{
		name:        name,
		streams:     streams,
		sync:        false,
		minLogLevel: 0,
	}
}

func (l *Logger) SetSynchronous(sync bool) {
	l.sync = sync
}

func (l *Logger) SetGlobalMinLogLevel(minLogLevel LogLevel) {
	l.minLogLevel = minLogLevel
}

func (l *Logger) AddStream(s StreamInterface) {
	l.streams = append(l.streams, s)
}

func (l *Logger) Log(e *LogEntry) {
	e.setLogger(l)

	if e.Level < l.minLogLevel {
		return
	}

	if l.sync {
		for _, stream := range l.streams {
			stream.Publish(e)
		}
	} else {
		for _, stream := range l.streams {
			go stream.Publish(e)
		}

	}
}

func (l *Logger) Logln(level LogLevel, message string) *LogEntry {
	e := NewLogEntry(level, message)

	l.Log(e)

	return e
}

func (l *Logger) LogF(level LogLevel, format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(level, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Trace(message string) *LogEntry {
	e := NewLogEntry(Trace, message)

	l.Log(e)

	return e
}

func (l *Logger) TraceF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Trace, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Debug(message string) *LogEntry {
	e := NewLogEntry(Debug, message)

	l.Log(e)

	return e
}

func (l *Logger) DebugF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Debug, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Info(message string) *LogEntry {
	e := NewLogEntry(Info, message)

	l.Log(e)

	return e
}

func (l *Logger) InfoF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Info, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Warn(message string) *LogEntry {
	e := NewLogEntry(Warn, message)

	l.Log(e)

	return e
}

func (l *Logger) WarnF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Warn, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Error(message string) *LogEntry {
	e := NewLogEntry(Error, message)

	l.Log(e)

	return e
}

func (l *Logger) ErrorF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Error, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

func (l *Logger) Fatal(message string) *LogEntry {
	e := NewLogEntry(Fatal, message)

	l.Log(e)

	return e
}

func (l *Logger) FatalF(format string, values ...interface{}) *LogEntry {
	e := NewLogEntry(Fatal, fmt.Sprintf(format, values...))

	l.Log(e)

	return e
}

// log.Logger compatibility

func (l *Logger) Println(val interface{}) {
	l.Logln(Info, fmt.Sprintf("%v", val))
}
