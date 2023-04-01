package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// GreaterThan matcher compares two values that are numeric or string values, and when
// called returns true if actual > expected. Strings are compared lexicographically with '>'.
// Returns a matcher that checks if actual is greater than expected.
func GreaterThan[A int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint16 | uint32 | uint64 | string](expected A) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("value greater than <%v>", expected)
	matcher.Matches = func(actual A) bool {
		return actual > expected
	}
	return matcher
}
