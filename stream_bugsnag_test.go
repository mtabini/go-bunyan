package bunyan

import (
	"github.com/bugsnag/bugsnag-go"
	"testing"
	"time"
)

func TestBugsnag(t *testing.T) {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "TestKey",
		ReleaseStage: "Test",
		Synchronous:  true,
	})

	s := NewBugsnagStream(Trace, nil)
	l := NewLogger("Test", []StreamInterface{s})

	l.Fatal("Achtung! Error! 1234")

	time.Sleep(time.Second)
}
