package bunyan

import (
	"encoding/json"
	"errors"
	"github.com/bugsnag/bugsnag-go"
)

type BugsnagStream struct {
	*Stream
}

func NewBugsnagStream(minLogLevel LogLevel, filter StreamFilter) (result *BugsnagStream) {
	return &BugsnagStream{
		Stream: &Stream{
			MinLogLevel: minLogLevel,
			Filter:      filter,
		},
	}
}

func (s *BugsnagStream) Publish(l *LogEntry) {
	if s.shouldPublish(l) {
		b, err := json.Marshal(l)

		if err != nil {
			return
		}

		data := map[string]interface{}{}

		if err := json.Unmarshal(b, &data); err != nil {
			return
		}

		err = errors.New(l.Error)

		if err == nil {
			err = errors.New(l.Message)
		}

		metadata := &bugsnag.MetaData{}
		metadata.AddStruct("Log", data)

		bugsnag.Notify(err, *metadata)
	}
}
