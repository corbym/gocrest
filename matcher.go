package gocrest

import (
	"fmt"
	"strings"
)

// Matcher provides the structure for matcher operations.
type Matcher[A any] struct {
	// Matches returns true if the function matches.
	Matches func(actual A) bool
	// Describe describes the matcher (e.g. "a value EqualTo(foo)" - foo being the expected value)
	Describe string
	// Actual is used by then.AssertThat if the matcher
	// needs to resolve the string description of the actual.
	// This is usually if the actual is a complex type.
	Actual string
	// ReasonString is a comment on why the matcher did not match, and set by the caller not the matcher.
	// Usually, this is set by the helper function, e.g. FooMatcher("foo").Reason("foo didn't foobar")
	ReasonString string
}

// Reason for the mismatch.
func (matcher *Matcher[A]) Reason(r string) *Matcher[A] {
	matcher.ReasonString = r
	return matcher
}

// Describes the Matcher to conform to the Stringer interface
func (matcher *Matcher[A]) String() string {
	return matcher.Describe
}

// Reasonf allows a formatted reason for the mismatch.
func (matcher *Matcher[A]) Reasonf(format string, args ...interface{}) *Matcher[A] {
	return matcher.Reason(fmt.Sprintf(format, args...))
}

// AppendActual appends an actual string to the matcher's actual description. This is useful if you want
// to preserve sub-matchers actual values. See is.AllOf() matcher for an example.
func (matcher *Matcher[A]) AppendActual(actualAsString string) {
	matcher.Actual += " " + actualAsString
	matcher.Actual = strings.TrimSpace(matcher.Actual)
}
