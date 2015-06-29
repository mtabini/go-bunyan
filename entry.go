package bunyan

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type Request struct {
	Method        string      `json:"method"`
	URL           string      `json:"url"`
	Headers       http.Header `json:"headers"`
	RemoteAddress string      `json:"remoteAddress"`
	Body          interface{} `json:"body,omitempty"`
}

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Headers    http.Header `json:"headers,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

type LogEntry struct {
	Data        interface{} `json:"data,omitempty"`
	Error       string      `json:"error,omitempty"`
	Hostname    string      `json:"hostname"`
	Level       LogLevel    `json:"level"`
	Message     string      `json:"msg"`
	Name        string      `json:"name"`
	ProcessID   int         `json:"pid"`
	Request     *Request    `json:"req,omitempty"`
	Response    *Response   `json:"res,omitempty"`
	StackTrace  string      `json:"trace,omitempty"`
	Time        time.Time   `json:"time"`
	CompletedIn string      `json:"completed_in,omitempty"`
	Version     int         `json:"v"`
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

func NewLogEntry(level LogLevel, message string) *LogEntry {
	result := logEntryTemplate

	result.Level = level
	result.Message = message
	result.Time = time.Now()
	result.Response = &Response{}

	return &result
}

func (l *LogEntry) SetData(data interface{}) {
	l.Data = data
}

func (l *LogEntry) SetRequest(r *http.Request) {
	l.Request = &Request{
		Method:        r.Method,
		URL:           r.URL.RequestURI(),
		Headers:       r.Header,
		RemoteAddress: r.RemoteAddr,
	}
}

func (l *LogEntry) SetRequestBody(body []byte) {
	if json.Unmarshal(body, &l.Request.Body) != nil {
		l.Request.Body = string(body)
	}
}

func (l *LogEntry) SetResponseStatusCode(statusCode int) {
	if statusCode > 0 {
		l.Response.StatusCode = statusCode
	} else {
		l.Response.StatusCode = http.StatusOK
	}
}

func (l *LogEntry) SetResponseBody(body []byte) {
	if json.Unmarshal(body, &l.Response.Body) != nil {
		l.Response.Body = string(body)
	}
}

func (l *LogEntry) SetResponseError(err error) {
	l.Error = err.Error()
}

func (l *LogEntry) SetCompletedIn(completedIn string) {
	l.CompletedIn = completedIn
}

func (l *LogEntry) SetStackTrace(trace string) {
	l.StackTrace = trace
}

func (l *LogEntry) setLogger(logger *Logger) {
	l.Name = logger.name
}

func (l *LogEntry) String() string {
	result, _ := json.Marshal(l)

	return string(result)
}
