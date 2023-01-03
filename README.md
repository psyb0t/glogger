# glogger

[![codecov](https://codecov.io/gh/psyb0t/glogger/branch/master/graph/badge.svg?token=QG0NA3QE7I)](https://codecov.io/gh/psyb0t/glogger)
[![goreportcard](https://goreportcard.com/badge/github.com/psyb0t/glogger)](https://goreportcard.com/report/github.com/psyb0t/glogger)
[![test](https://github.com/psyb0t/glogger/actions/workflows/test.yml/badge.svg)](https://github.com/psyb0t/glogger/actions/workflows/test.yml)
[![golangci-lint](https://github.com/psyb0t/glogger/actions/workflows/lint.yml/badge.svg)](https://github.com/psyb0t/glogger/actions/workflows/lint.yml)

glogger is a Go package that provides a logger that can be used to write log messages to an output destination, such as the standard output or a file. It has various levels of log severity (e.g. debug, info, warning, error, fatal) and supports two formats: inline-string and JSON.

## Installation

To install glogger, run:

`go get github.com/psyb0t/glogger`

## Usage

The global variables `output`, `logFormat`, and `logLevel` control the output destination, format, and minimum severity level of log messages, respectively. You can set these variables using the functions `SetOutput`, `SetLogFormat`, and `SetLogLevel`.

To create a new logger, use the `New` function and pass in a `Caller` struct that contains the caller information (e.g. service, package, receiver, function). The `Logger` interface provides methods for logging messages at different severity levels: `Debug`, `Info`, `Warn`, `Error`, and `Fatal`.

The `Err`, `RequestID`, `TraceID`, `SpanID`, and `Data` methods can be used to add extra data to the log entry. The `Err` method sets the error field in the log entry to the given value. The `RequestID`, `TraceID`, and `SpanID` methods set the corresponding fields in the log entry to the given values. The `Data` method adds a data field to the log entry as key => val.

## Example

```golang
package main

import (
	"errors"
	"fmt"
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
	logger := glogger.New(glogger.Caller{Service: "my-service", Package: "main", Function: "main"})

	// Log messages at different severity levels with various extra data.
	logger.RequestID("abc123").Debug("hello world")
	logger.TraceID("zzzsss").Info("hello world")
	logger.TraceID("zzzsss").SpanID("123456").Warn("hello world")
	logger.Err(errors.New("something bad happened")).Error("hello world")
	logger.Data("key", "value").Fatal("hello world")
}
```

The log file `app.log` will contain the following messages:

```
{"service":"my-service","package":"main","function":"main","time":"2023-01-04T00:50:04+02:00","level":"INFO","message":"hello world","requestID":"abc123","traceID":"zzzsss"}
{"service":"my-service","package":"main","function":"main","time":"2023-01-04T00:50:04+02:00","level":"WARN","message":"hello world","traceID":"zzzsss","spanID":"123456"}
{"service":"my-service","package":"main","function":"main","time":"2023-01-04T00:50:04+02:00","level":"ERROR","message":"hello world","error":"something bad happened"}
{"service":"my-service","package":"main","function":"main","time":"2023-01-04T00:50:04+02:00","level":"FATAL","message":"hello world","data":{"key":"value"}}
```
