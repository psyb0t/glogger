package glogger

// Caller is the source of the logs. It is included in the log data
type Caller struct {
	Service  string `json:"service,omitempty"`
	Package  string `json:"package,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Function string `json:"function,omitempty"`
}

// String returns the string representation of a Caller
func (t Caller) String() string {
	isEmptyStrings := func(ss ...string) bool {
		for _, s := range ss {
			if s != "" {
				return false
			}
		}

		return true
	}

	s := ""

	if t.Service != "" {
		s += t.Service

		if !(isEmptyStrings(t.Package, t.Receiver, t.Function)) {
			s += "|"
		}
	}

	if t.Package != "" {
		s += t.Package

		if !(isEmptyStrings(t.Receiver, t.Function)) {
			s += ":"
		}
	}

	if t.Receiver != "" {
		s += t.Receiver

		if t.Function != "" {
			s += "."
		}
	}

	if t.Function != "" {
		s += t.Function
	}

	return s
}
