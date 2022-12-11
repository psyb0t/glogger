package glogger

import (
	"fmt"
	"io"
	"os"
)

// LogFormat represents the format of a log message
type LogFormat uint8

const (
	// LogFormatDefault is the default log format
	LogFormatDefault LogFormat = iota
	// LogFormatJSON is the JSON log format
	LogFormatJSON
)

// LogLevel is the type that represents the log level
type LogLevel uint8

const (
	// LogLevelDebug is the debug log level
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is the info log level
	LogLevelInfo
	// LogLevelWarn is the warning log level
	LogLevelWarn
	// LogLevelError is the error log level
	LogLevelError
	// LogLevelFatal is the fatal log level
	LogLevelFatal
)

var logLevelString = map[LogLevel]string{
	LogLevelDebug: "DEBUG",
	LogLevelInfo:  "INFO",
	LogLevelWarn:  "WARN",
	LogLevelError: "ERROR",
	LogLevelFatal: "FATAL",
}

var (
	output    io.Writer = os.Stdout
	logFormat LogFormat = LogFormatDefault
	logLevel  LogLevel  = LogLevelDebug
)

// SetLogFormat sets the global variable that controls
// the format of the log lines being printed
func SetLogFormat(format LogFormat) {
	logFormat = format
}

// SetLogLevel sets the global variable that controls
// the minimum severity level required for a message to
// pass through
func SetLogLevel(level LogLevel) {
	logLevel = level
}

// SetOutput sets the global variable that specifies the output writer
// where log messages will be written.
//
// For example, the following code configures the logger to write messages
// to a file named "app.log":
//
//     f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//     if err != nil {
//         log.Fatalf("error opening file: %v", err)
//     }
//     defer f.Close()
//
//     glogger.SetOutput(f)
//
func SetOutput(o io.Writer) {
	output = o
}

func _print(msg string) {
	fmt.Fprintln(output, msg)
}