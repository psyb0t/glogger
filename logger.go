package glogger

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
