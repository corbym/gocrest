package gocrest

import "fmt"

//Matcher provides the structure for matcher operations.
type Matcher struct {
	// Matches returns true if the function matches.
	Matches func(actual interface{}) bool
	// Describe describes the matcher (e.g. "a value EqualTo(foo)"
	Describe string
	// Actual is used if the matcher needs to resolve the string description of the matcher.
	// This is usually if the actual is a complex type.
	Actual string
	// ReasonString is a comment on why the matcher did not match.
	ReasonString string
}

//Reason for the mismatch.
func (matcher *Matcher) Reason(r string) *Matcher {
	matcher.ReasonString = r
	return matcher
}

//Reasonf allows a formatted reason for the mismatch.
func (matcher *Matcher) Reasonf(format string, args ...interface{}) *Matcher {
	return matcher.Reason(fmt.Sprintf(format, args...))
}
