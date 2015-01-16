package bunyan

import (
	"testing"
)

func TestLogEntryCopy(t *testing.T) {
	s := NewStdoutStream(Info, nil)
	l := NewLogger("test", []StreamInterface{s})

	e := NewLogEntry(Warn, "Test")

	e2 := NewLogEntry(Warn, "Test 2")

	l.Log(e)

	if e.Message == e2.Message {
		t.Error("Creating a new log entry results in a pointer to the template.")
		t.Fail()
	}
}
