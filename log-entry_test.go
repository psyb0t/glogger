package glogger

import (
	"bytes"
	"errors"
	"strings"
	"sync"
	"testing"
)

type mutexBuffer struct {
	b bytes.Buffer
	m sync.Mutex
}

func (b *mutexBuffer) Read(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Read(p)
}
func (b *mutexBuffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Write(p)
}
func (b *mutexBuffer) String() string {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.String()
}

func TestLogEntry(t *testing.T) {
	t.Run("test data and metadata setters", func(t *testing.T) {
		testDataKey := "key"

		tests := []struct {
			name       string
			method     func(Logger) Logger
			methodName string
			expected   string
		}{
			{
				name:       "Err",
				method:     func(l Logger) Logger { return l.Err(errors.New("error message")) },
				methodName: "Err",
				expected:   "error message",
			},
			{
				name:       "RequestID",
				method:     func(l Logger) Logger { return l.RequestID("request-id") },
				methodName: "RequestID",
				expected:   "request-id",
			},
			{
				name:       "TraceID",
				method:     func(l Logger) Logger { return l.TraceID("trace-id") },
				methodName: "TraceID",
				expected:   "trace-id",
			},
			{
				name:       "SpanID",
				method:     func(l Logger) Logger { return l.SpanID("span-id") },
				methodName: "SpanID",
				expected:   "span-id",
			},
			{
				name:       "Data",
				method:     func(l Logger) Logger { return l.Data(testDataKey, "value") },
				methodName: "Data",
				expected:   "value",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				var l Logger
				l = logEntry{}
				l = test.method(l)

				entry := l.(logEntry)

				actual := ""
				switch test.methodName {
				case "Err":
					actual = entry.E
				case "RequestID":
					actual = entry.ReqID
				case "TraceID":
					actual = entry.TrcID
				case "SpanID":
					actual = entry.SpnID
				case "Data":
					var ok bool
					var val interface{}

					if val, ok = entry.D[testDataKey]; !ok {
						t.Errorf("Expected key %s not set in the data field", testDataKey)
					}

					if actual, ok = val.(string); !ok {
						t.Errorf("Data field key value is not a string as expected %v", val)
					}
				}

				if actual != test.expected {
					t.Errorf("Expected %s but got %s on method %s", test.expected, actual, test.methodName)
				}
			})
		}

	})

	t.Run("test print methods", func(t *testing.T) {
		l := logEntry{}

		tests := []struct {
			name     string
			logFunc  func(msg string)
			input    string
			expected string
		}{
			{
				name:     "Debug",
				logFunc:  l.Debug,
				input:    "debug msg",
				expected: "debug msg",
			},
			{
				name:     "Info",
				logFunc:  l.Info,
				input:    "info msg",
				expected: "info msg",
			},
			{
				name:     "Warn",
				logFunc:  l.Warn,
				input:    "warn msg",
				expected: "warn msg",
			},
			{
				name:     "Error",
				logFunc:  l.Error,
				input:    "error msg",
				expected: "error msg",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				SetLogFormat(LogFormatInlineString)
				SetLogLevel(LogLevelDebug)

				buf := &mutexBuffer{}
				SetOutput(buf)

				test.logFunc(test.input)

				bufStr := buf.String()
				if !strings.Contains(bufStr, test.expected) {
					t.Errorf("expected output to contain %q, got %q", test.expected, bufStr)
				}
			})
		}
	})

	t.Run("test race condition", func(t *testing.T) {
		l := logEntry{}
		buf := &mutexBuffer{}
		SetOutput(buf)

		var wg sync.WaitGroup

		for i := 0; i < 50; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				l.Data("key", "val").Info("info msg")
			}()

			wg.Add(1)
			go func() {
				defer wg.Done()

				l.Err(errors.New("error")).Error("error msg")
			}()
		}

		wg.Wait()
	})
}
