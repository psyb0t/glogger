package glogger

// Caller is the source of the logs. It is included in the log data
type Caller struct {
	Package  string `json:"package,omitempty"`
	Receiver string `json:"receiver,omitempty"`
	Function string `json:"function,omitempty"`
}

// String returns the string representation of a Caller
func (t Caller) String() string {
	s := ""

	if t.Package != "" {
		s = t.Package

		if t.Receiver != "" || t.Function != "" {
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
