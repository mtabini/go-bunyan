package bunyan

import (
	"testing"
)

func TestStdoutStream(t *testing.T) {
	s := NewStdoutStream(Info, nil)

	l := NewLogger("Test", []StreamInterface{s})

	l.Fatal("Test")
}
