package glogger

import (
	"bytes"
	"strings"
	"testing"
)

func TestStrToLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected LogLevel
	}{
		{"DEBUG", LogLevelDebug},
		{"Info", LogLevelInfo},
		{"warn", LogLevelWarn},
		{"error", LogLevelError},
		{"fatal", LogLevelFatal},
		{"invalid", LogLevelDebug},
	}

	for _, test := range tests {
		result := StrToLogLevel(test.input)
		if result != test.expected {
			t.Errorf("StrToLogLevel(%q) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestSetLogFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   LogFormat
		expected LogFormat
	}{
		{
			name:     "default format",
			format:   LogFormatInlineString,
			expected: LogFormatInlineString,
		},
		{
			name:     "JSON format",
			format:   LogFormatJSON,
			expected: LogFormatJSON,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetLogFormat(test.format)
			if logFormat != test.expected {
				t.Errorf("expected log format %d, got %d", test.expected, logFormat)
			}
		})
	}
}

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		level    LogLevel
		expected LogLevel
	}{
		{
			name:     "debug level",
			level:    LogLevelDebug,
			expected: LogLevelDebug,
		},
		{
			name:     "info level",
			level:    LogLevelInfo,
			expected: LogLevelInfo,
		},
		{
			name:     "warn level",
			level:    LogLevelWarn,
			expected: LogLevelWarn,
		},
		{
			name:     "error level",
			level:    LogLevelError,
			expected: LogLevelError,
		},
		{
			name:     "fatal level",
			level:    LogLevelFatal,
			expected: LogLevelFatal,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SetLogLevel(test.level)
			if logLevel != test.expected {
				t.Errorf("expected log level %d, got %d", test.expected, logLevel)
			}
		})
	}
}

func TestSetOutput(t *testing.T) {
	buf := &bytes.Buffer{}
	SetOutput(buf)

	msg := "this is a message"
	_print(msg)

	result := strings.TrimRight(buf.String(), "\n")
	if result != msg {
		t.Errorf("unexpected log output: got %q, want %q", result, msg)
	}
}
