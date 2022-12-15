package glogger

import "testing"

func TestCaller_String(t *testing.T) {
	tests := []struct {
		name     string
		caller   Caller
		expected string
	}{
		{
			name: "all fields empty",
			caller: Caller{
				Service:  "",
				Package:  "",
				Receiver: "",
				Function: "",
			},
			expected: "",
		},
		{
			name: "service only",
			caller: Caller{
				Service:  "my-service",
				Package:  "",
				Receiver: "",
				Function: "",
			},
			expected: "my-service",
		},
		{
			name: "package only",
			caller: Caller{
				Service:  "",
				Package:  "my-package",
				Receiver: "",
				Function: "",
			},
			expected: "my-package",
		},
		{
			name: "receiver only",
			caller: Caller{
				Service:  "",
				Package:  "",
				Receiver: "my-receiver",
				Function: "",
			},
			expected: "my-receiver",
		},
		{
			name: "function only",
			caller: Caller{
				Service:  "",
				Package:  "",
				Receiver: "",
				Function: "my-function",
			},
			expected: "my-function",
		},
		{
			name: "service and package only",
			caller: Caller{
				Service:  "my-service",
				Package:  "my-package",
				Receiver: "",
				Function: "",
			},
			expected: "my-service|my-package",
		},
		{
			name: "service and receiver only",
			caller: Caller{
				Service:  "my-service",
				Package:  "",
				Receiver: "my-receiver",
				Function: "",
			},
			expected: "my-service|my-receiver",
		},
		{
			name: "service and function only",
			caller: Caller{
				Service:  "my-service",
				Package:  "",
				Receiver: "",
				Function: "my-function",
			},
			expected: "my-service|my-function",
		},
		{
			name: "package and receiver only",
			caller: Caller{
				Service:  "",
				Package:  "my-package",
				Receiver: "my-receiver",
				Function: "",
			},
			expected: "my-package:my-receiver",
		},
		{
			name: "package and function only",
			caller: Caller{
				Service:  "",
				Package:  "my-package",
				Receiver: "",
				Function: "my-function",
			},
			expected: "my-package:my-function",
		},
		{
			name: "receiver and function only",
			caller: Caller{
				Service:  "",
				Package:  "",
				Receiver: "my-receiver",
				Function: "my-function",
			},
			expected: "my-receiver.my-function",
		},
		{
			name: "all fields set",
			caller: Caller{
				Service:  "my-service",
				Package:  "my-package",
				Receiver: "my-receiver",
				Function: "my-function",
			},
			expected: "my-service|my-package:my-receiver.my-function",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.caller.String()
			if result != test.expected {
				t.Errorf("unexpected result: got %q, want %q", result, test.expected)
			}
		})
	}
}
