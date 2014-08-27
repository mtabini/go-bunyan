package bunyan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func validateLogEntry(t *testing.T, newEntry *LogEntry, originalEntry *LogEntry) {
	if *newEntry != *originalEntry {
		t.Error("Two entries are not equal!")
		t.Errorf("Original; %s", originalEntry)
		t.Errorf("New: %s", newEntry)

		t.Fail()
	}
}

func TestFileStream(t *testing.T) {
	f, err := ioutil.TempFile("", "bunyan")

	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
		t.Fail()
		return
	}

	path := f.Name()

	f.Close()

	defer os.Remove(path)

	s, err := NewFileStream(Info, nil, path)

	if err != nil {
		t.Fatalf("Error creating file stream: %v", err)
		t.Fail()
		return
	}

	l := NewLogger("Test", []StreamInterface{s})

	originalEntry := l.Fatal("Test")

	time.Sleep(1000000)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
		t.Fail()
		return
	}

	newEntry := &LogEntry{}

	err = json.Unmarshal(data, newEntry)

	if err != nil {
		t.Fatalf("Error unmarshaling test file contents: %v", err)
		t.Fail()
		return
	}

	validateLogEntry(t, newEntry, originalEntry)
}
