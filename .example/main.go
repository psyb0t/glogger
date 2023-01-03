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

/*
{"time":"2023-01-03T00:00:00Z","level":"DEBUG","message":"hello world","requestID":"abc123"}
{"time":"2023-01-03T00:00:00Z","level":"INFO","message":"hello world","traceID":"zzzsss"}
{"time":"2023-01-03T00:00:00Z","level":"WARN","message":"hello world","traceID":"zzzsss","spanID":"123456"}
{"time":"2023-01-03T00:00:00Z","level":"ERROR","message":"hello world","error":"something bad happened"}
{"time":"2023-01-03T00:00:00Z","level":"FATAL","message":"hello world","data":{"key":"value"}}
*/
