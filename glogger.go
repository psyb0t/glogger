// Package glogger provides a logger that can be used to write log messages
// to an output destination, such as the standard output or a file. It has
// various levels of log severity (e.g. debug, info, warning, error, fatal)
// and the following formats: inline-string, json.
//
// The global variables output, logFormat, and logLevel control the output
// destination, format, and minimum severity level of log messages,
// respectively. You can set these variables using the functions SetOutput,
// SetLogFormat, and SetLogLevel.
//
// The function _print writes the given message to the output destination.
// The functions Debug, Info, Warn, Error, and Fatal write log messages
// with the corresponding severity levels, if the log level of the message
// is equal to or higher than the minimum log level set.
package glogger

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// LogLevel is the type that represents the log level
type LogLevel uint8

// Constants representing the available log levels.
const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// logLevelString is a map of LogLevel constants to their string representations.
var logLevelString = map[LogLevel]string{
	LogLevelDebug: "DEBUG",
	LogLevelInfo:  "INFO",
	LogLevelWarn:  "WARN",
	LogLevelError: "ERROR",
	LogLevelFatal: "FATAL",
}

// StrToLogLevel matches the given string to a supported LogLevel.
// If no matches are found it returns LogLevelDebug
func StrToLogLevel(str string) LogLevel {
	for k, v := range logLevelString {
		if strings.EqualFold(str, v) {
			return k
		}
	}

	return LogLevelDebug
}

// LogFormat represents the format of a log message
type LogFormat uint8

// Constants representing the available log formats.
const (
	LogFormatInlineString LogFormat = iota // (default)
	LogFormatJSON
)

// logFormatString is a map of LogFormat constants to their string representations.
var logFormatString = map[LogFormat]string{
	LogFormatInlineString: "inline-string",
	LogFormatJSON:         "json",
}

// StrToLogFormat matches the given string to a supported LogFormat.
// If no matches are found it returns LogFormatInlineString
func StrToLogFormat(str string) LogFormat {
	for k, v := range logFormatString {
		if strings.EqualFold(str, v) {
			return k
		}
	}

	return LogFormatInlineString
}

var (
	output    io.Writer = os.Stdout
	logFormat LogFormat = LogFormatInlineString
	logLevel  LogLevel  = LogLevelDebug
)

// SetLogFormat sets the global variable that controls
// the format of the log lines being printed
func SetLogFormat(format LogFormat) {
	logFormat = format
}

// SetLogLevel sets the global variable that controls
// the minimum severity level required for a message to
// pass through. If an invalid value is passed
// it sets it to LogLevelDebug
func SetLogLevel(level LogLevel) {
	if level > LogLevelFatal {
		logLevel = LogLevelDebug

		return
	}

	logLevel = level
}

// SetOutput sets the global variable that specifies the output writer
// where log messages will be written.
//
// For example, the following code configures the logger to write messages
// to a file named "app.log":
//
//	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//	    log.Fatalf("error opening file: %v", err)
//	}
//	defer f.Close()
//
//	glogger.SetOutput(f)
func SetOutput(o io.Writer) {
	output = o
}

func _print(msg string) {
	fmt.Fprintln(output, msg)
}
