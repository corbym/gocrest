package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// LessThan matcher compares two values that are numeric or string values, and when
// called returns true if actual < expected. Strings are compared lexicographically with '<'.
// The matcher will always return false for unknown types.
// Actual and expected types must be the same underlying type, or the function will panic.
// Returns a matcher that checks if actual is greater than expected.
func LessThan[A int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint16 | uint32 | uint64 | string](expected A) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("value less than <%v>", expected)
	matcher.Matches = func(actual A) bool {
		return actual < expected
	}
	return matcher
}
