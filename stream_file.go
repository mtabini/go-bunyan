package bunyan

import (
	"os"
)

type FileStream struct {
	*Stream
	outputFile *os.File
}

func NewFileStream(minLogLevel LogLevel, filter StreamFilter, path string) (result *FileStream, err error) {
	outputFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		return
	}

	result = &FileStream{
		Stream: &Stream{
			MinLogLevel: minLogLevel,
			Filter:      filter,
		},
		outputFile: outputFile,
	}

	return
}

func (s *FileStream) Publish(l *LogEntry) {
	if s.shouldPublish(l) {
		s.outputFile.WriteString(l.String())
		s.outputFile.WriteString("\n")
	}
}
