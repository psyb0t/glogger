package glogger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Logger is the interface that defines methods for logging messages at different severity levels.
type Logger interface {
	// Err sets the error field in the log entry to the given value.
	Err(err error) Logger
	// RequestID sets the requestID field in the log entry to the given value.
	RequestID(id string) Logger
	// TraceID sets the traceID field in the log entry to the given value.
	TraceID(id string) Logger
	// SpanID sets the spanID field in the log entry to the given value.
	SpanID(id string) Logger
	// Data adds a data field in the log entry as key => val.
	Data(key string, val interface{}) Logger

	// Debug logs a message at the debug severity level.
	Debug(msg string)
	// Info logs a message at the info severity level.
	Info(msg string)
	// Warn logs a message at the warning severity level.
	Warn(msg string)
	// Error logs a message at the error severity level.
	Error(msg string)
	// Fatal logs a message at the fatal severity level and terminates the program.
	Fatal(msg string)
}

type logEntry struct {
	Caller
	Time      string                 `json:"time,omitempty"`
	Level     string                 `json:"level,omitempty"`
	Message   string                 `json:"message,omitempty"`
	Error     string                 `json:"error,omitempty"`
	RequestID string                 `json:"requestID,omitempty"`
	TraceID   string                 `json:"traceID,omitempty"`
	SpanID    string                 `json:"spanID,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type logger struct {
	caller Caller
	entry  logEntry
}

// New creates and returns a new Logger with the given caller.
func New(caller Caller) Logger {
	return &logger{caller: caller}
}

func (l *logger) Err(err error) Logger {
	l.entry.Error = err.Error()

	return l
}

func (l *logger) RequestID(id string) Logger {
	l.entry.RequestID = id

	return l
}

func (l *logger) TraceID(id string) Logger {
	l.entry.TraceID = id

	return l
}

func (l *logger) SpanID(id string) Logger {
	l.entry.SpanID = id

	return l
}

func (l *logger) Data(key string, val interface{}) Logger {
	if l.entry.Data == nil {
		l.entry.Data = map[string]interface{}{}
	}

	l.entry.Data[key] = val

	return l
}

func (l *logger) Debug(msg string) {
	if !(logLevel < LogLevelDebug+1) {
		return
	}

	l.print(LogLevelDebug, msg)
}

func (l *logger) Info(msg string) {
	if !(logLevel < LogLevelInfo+1) {
		return
	}

	l.print(LogLevelInfo, msg)
}

func (l *logger) Warn(msg string) {
	if !(logLevel < LogLevelWarn+1) {
		return
	}

	l.print(LogLevelWarn, msg)
}

func (l *logger) Error(msg string) {
	if !(logLevel < LogLevelError+1) {
		return
	}

	l.print(LogLevelError, msg)
}

func (l *logger) Fatal(msg string) {
	l.print(LogLevelFatal, msg)

	os.Exit(1)
}

func (l *logger) print(lvl LogLevel, msg string) {
	defer func() {
		l.entry = logEntry{}
	}()

	l.entry.Caller = l.caller
	l.entry.Time = time.Now().Format(time.RFC3339)
	l.entry.Level = logLevelString[lvl]
	l.entry.Message = msg

	if logFormat == LogFormatJSON {
		lJSON, err := json.Marshal(l.entry)
		if err != nil {
			panic(err)
		}

		_print(string(lJSON))

		return
	}

	output := fmt.Sprintf("%s %s %s", l.entry.Time, l.entry.Level, l.caller.String())

	if l.entry.RequestID != "" {
		output += fmt.Sprintf(" RequestID: %s |", l.entry.RequestID)
	}

	if l.entry.TraceID != "" {
		output += fmt.Sprintf(" TraceID: %s |", l.entry.TraceID)
	}

	if l.entry.SpanID != "" {
		output += fmt.Sprintf(" SpanID: %s |", l.entry.SpanID)
	}

	if len(l.entry.Data) > 0 {
		output += fmt.Sprintf(" Data: %v |", l.entry.Data)
	}

	if l.entry.Error != "" {
		output += fmt.Sprintf(" Error: %s |", l.entry.Error)
	}

	_print(fmt.Sprintf("%s %s", output, l.entry.Message))
}
