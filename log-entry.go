package glogger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type logEntry struct {
	Caller
	Time    string                 `json:"time,omitempty"`
	Level   string                 `json:"level,omitempty"`
	Message string                 `json:"message,omitempty"`
	E       string                 `json:"error,omitempty"`
	ReqID   string                 `json:"requestID,omitempty"`
	TrcID   string                 `json:"traceID,omitempty"`
	SpnID   string                 `json:"spanID,omitempty"`
	D       map[string]interface{} `json:"data,omitempty"`
}

// New creates and returns a new Logger with the given caller.
func New(caller Caller) Logger {
	return logEntry{Caller: caller}
}

func (l logEntry) Err(err error) Logger {
	l.E = err.Error()

	return l
}

func (l logEntry) RequestID(id string) Logger {
	l.ReqID = id

	return l
}

func (l logEntry) TraceID(id string) Logger {
	l.TrcID = id

	return l
}

func (l logEntry) SpanID(id string) Logger {
	l.SpnID = id

	return l
}

func (l logEntry) Data(key string, val interface{}) Logger {
	if l.D == nil {
		l.D = map[string]interface{}{}
	}

	l.D[key] = val

	return l
}

func (l logEntry) Debug(msg string) {
	if !(logLevel < LogLevelDebug+1) {
		return
	}

	l.print(LogLevelDebug, msg)
}

func (l logEntry) Info(msg string) {
	if !(logLevel < LogLevelInfo+1) {
		return
	}

	l.print(LogLevelInfo, msg)
}

func (l logEntry) Warn(msg string) {
	if !(logLevel < LogLevelWarn+1) {
		return
	}

	l.print(LogLevelWarn, msg)
}

func (l logEntry) Error(msg string) {
	if !(logLevel < LogLevelError+1) {
		return
	}

	l.print(LogLevelError, msg)
}

func (l logEntry) Fatal(msg string) {
	l.print(LogLevelFatal, msg)

	os.Exit(1)
}

func (l logEntry) print(lvl LogLevel, msg string) {
	l.Time = time.Now().Format(time.RFC3339)
	l.Level = logLevelString[lvl]
	l.Message = msg

	if logFormat == LogFormatJSON {
		lJSON, err := json.Marshal(l)
		if err != nil {
			panic(err)
		}

		_print(string(lJSON))

		return
	}

	output := fmt.Sprintf("%s %s %s", l.Time, l.Level, l.Caller.String())

	if l.ReqID != "" {
		output += fmt.Sprintf(" RequestID: %s |", l.ReqID)
	}

	if l.TrcID != "" {
		output += fmt.Sprintf(" TraceID: %s |", l.TrcID)
	}

	if l.SpnID != "" {
		output += fmt.Sprintf(" SpanID: %s |", l.SpnID)
	}

	if len(l.D) > 0 {
		output += fmt.Sprintf(" Data: %v |", l.D)
	}

	if l.E != "" {
		output += fmt.Sprintf(" Error: %s |", l.E)
	}

	_print(fmt.Sprintf("%s %s", output, l.Message))
}
