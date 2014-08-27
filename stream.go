package bunyan

type StreamInterface interface {
	Publish(*LogEntry)
}

type StreamFilter func(*LogEntry) bool

type Stream struct {
	MinLogLevel LogLevel
	Filter      StreamFilter
}

func (s *Stream) shouldPublish(l *LogEntry) bool {
	if l.Level < s.MinLogLevel {
		return false
	}

	if s.Filter != nil {
		return s.Filter(l)
	}

	return true
}
