# glogger

[![codecov](https://codecov.io/gh/psyb0t/glogger/branch/master/graph/badge.svg?token=QG0NA3QE7I)](https://codecov.io/gh/psyb0t/glogger)
[![goreportcard](https://goreportcard.com/badge/github.com/psyb0t/glogger)](https://goreportcard.com/report/github.com/psyb0t/glogger)
[![test](https://github.com/psyb0t/glogger/actions/workflows/test.yml/badge.svg)](https://github.com/psyb0t/glogger/actions/workflows/test.yml)
[![golangci-lint](https://github.com/psyb0t/glogger/actions/workflows/lint.yml/badge.svg)](https://github.com/psyb0t/glogger/actions/workflows/lint.yml)

This package provides a simple logging library for Go programs.

## Installation

To install glogger, run:

`go get github.com/psyb0t/glogger`

## Usage

The glogger package defines the `Logger` interface, which has the following methods:

```golang
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
```

To use the logger, you need to create an instance of the `logger` struct, which implements the `Logger` interface. For example:

```golang
logger := glogger.New(glogger.Caller{Package: "github.com/user/pkg", Receiver: "foo", Function: "bar"})
```

The `logger` struct has the following fields:

- `caller`: the `Caller` object that provides information about the caller of the logger method (service, package, receiver, and function name).

Once you have created an instance of the logger struct, you can use its methods to log messages at different severity levels:

```golang
logger.Debug("this is a debug message")
logger.Info("this is an info message")
logger.Warn("this is a warning message")
logger.Error("this is an error message")
logger.Fatal("this is a fatal message")
```

The `glogger` package also provides a few global functions to customize the behavior of the logger:

```golang
// StrToLogLevel matches the given string to a supported LogLevel.
// If no matches are found it returns LogLevelDebug
func StrToLogLevel(str string) LogLevel

// SetLogFormat sets the global variable that controls
// the format of the log lines being printed
func SetLogFormat(format LogFormat)

// SetLogLevel sets the global variable that controls
// the minimum severity level required for a message to
// pass through
func SetLogLevel(level LogLevel)

// SetOutput sets the global variable that specifies the output writer
// where log messages will be written.
func SetOutput(o io.Writer)
```

The `SetLogFormat` function allows you to specify the format of the log lines being printed (default or JSON). The `SetLogLevel` function allows you to specify the minimum severity level required for a message to pass through (e.g., only log messages with severity level `Error` or higher). And the `SetOutput` function allows you to specify the output writer where log messages will be written (e.g., a file, the standard output, etc.).

## Example

Here is a complete example of how to use the `glogger` package to log messages to a file named `app.log`:

```golang
package main

import (
	"os"

	"github.com/psyb0t/glogger"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	defer f.Close()

	// Set the output writer where log messages will be written.
	glogger.SetOutput(f)

	// Set the log format to JSON.
	glogger.SetLogFormat(glogger.LogFormatJSON)

	// Set the log level to "Info" or higher.
	glogger.SetLogLevel(glogger.LogLevelInfo)

	// Create a logger instance.
	logger := glogger.New(glogger.Caller{Package: "github.com/user/pkg", Receiver: "foo", Function: "bar"})

	// Log messages at different severity levels.
	logger.Debug("hello world")
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	logger.Fatal("hello world")
}
```

The log file `app.log` will contain the following messages:

```
{"time":"2022-12-11T17:19:49+01:00","logLevel":"INFO","caller":"github.com/user/pkg:foo.bar","value":["hello world"]}
{"time":"2022-12-11T17:19:49+01:00","logLevel":"WARN","caller":"github.com/user/pkg:foo.bar","value":["hello world"]}
{"time":"2022-12-11T17:19:49+01:00","logLevel":"ERROR","caller":"github.com/user/pkg:foo.bar","value":["hello world"]}
{"time":"2022-12-11T17:19:49+01:00","logLevel":"FATAL","caller":"github.com/user/pkg:foo.bar","value":["hello world"]}
```
