package glogger

import "testing"

func TestCaller_String(t *testing.T) {
	tests := []struct {
		caller   Caller
		expected string
	}{
		{Caller{}, ""},
		{Caller{Package: "github.com/user/pkg"}, "github.com/user/pkg"},
		{Caller{Receiver: "foo"}, "foo"},
		{Caller{Function: "bar"}, "bar"},
		{Caller{Package: "github.com/user/pkg", Receiver: "foo"}, "github.com/user/pkg:foo"},
		{Caller{Package: "github.com/user/pkg", Function: "bar"}, "github.com/user/pkg:bar"},
		{Caller{Receiver: "foo", Function: "bar"}, "foo.bar"},
		{Caller{Package: "github.com/user/pkg", Receiver: "foo", Function: "bar"}, "github.com/user/pkg:foo.bar"},
	}

	for _, test := range tests {
		actual := test.caller.String()
		if actual != test.expected {
			t.Errorf("Expected %q but got %q", test.expected, actual)
		}
	}
}
