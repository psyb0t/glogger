package glogger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New(Caller{})

	tests := []struct {
		name     string
		logFunc  func(...interface{})
		input    []interface{}
		expected string
	}{
		{
			name:     "Debug",
			logFunc:  logger.Debug,
			input:    []interface{}{"debug msg"},
			expected: "debug msg",
		},
		{
			name:     "Info",
			logFunc:  logger.Info,
			input:    []interface{}{"info msg"},
			expected: "info msg",
		},
		{
			name:     "Warn",
			logFunc:  logger.Warn,
			input:    []interface{}{"warn msg"},
			expected: "warn msg",
		},
		{
			name:     "Error",
			logFunc:  logger.Error,
			input:    []interface{}{"error msg"},
			expected: "error msg",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetLogFormat(LogFormatDefault)
			SetLogLevel(LogLevelDebug)

			buf := &bytes.Buffer{}
			SetOutput(buf)

			test.logFunc(test.input...)

			bufStr := buf.String()
			if !strings.Contains(bufStr, test.expected) {
				t.Errorf("expected output to contain %q, got %q", test.expected, bufStr)
			}
		})
	}
}
