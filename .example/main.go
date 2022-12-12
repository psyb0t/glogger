package main

import (
	"fmt"
	"os"

	"github.com/psyb0t/glogger"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
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
