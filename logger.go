package glogger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Logger is the interface that defines methods for logging messages at different severity levels.
type Logger interface {
	// Debug logs a message at the debug severity level.
	Debug(v ...interface{})

	// Info logs a message at the info severity level.
	Info(v ...interface{})

	// Warn logs a message at the warning severity level.
	Warn(v ...interface{})

	// Error logs a message at the error severity level.
	Error(v ...interface{})

	// Fatal logs a message at the fatal severity level and terminates the program.
	Fatal(v ...interface{})
}

// New creates and returns a new Logger with the given caller.
func New(caller Caller) Logger {
	return logger{caller: caller}
}

type logger struct {
	caller Caller
}

func (l logger) Debug(v ...interface{}) {
	if !(logLevel < LogLevelDebug+1) {
		return
	}

	l.print(LogLevelDebug, v...)
}

func (l logger) Info(v ...interface{}) {
	if !(logLevel < LogLevelInfo+1) {
		return
	}

	l.print(LogLevelInfo, v...)
}

func (l logger) Warn(v ...interface{}) {
	if !(logLevel < LogLevelWarn+1) {
		return
	}

	l.print(LogLevelWarn, v...)
}

func (l logger) Error(v ...interface{}) {
	if !(logLevel < LogLevelError+1) {
		return
	}

	l.print(LogLevelError, v...)
}

func (l logger) Fatal(v ...interface{}) {
	l.print(LogLevelFatal, v...)

	os.Exit(1)
}

func (l logger) print(lvl LogLevel, v ...interface{}) {
	d := map[string]interface{}{
		"time":     time.Now().Format(time.RFC3339),
		"logLevel": logLevelString[lvl],
		"caller":   l.caller.String(),
		"value":    v,
	}

	if logFormat == LogFormatJSON {
		dJSON, err := json.Marshal(d)
		if err != nil {
			panic(err)
		}

		_print(string(dJSON))

		return
	}

	_print(strings.TrimRight(fmt.Sprintln(append([]interface{}{
		d["time"], d["logLevel"], d["caller"],
	}, v...)...), "\n"))
}
