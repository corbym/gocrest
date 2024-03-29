package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// LessThan matcher compares two values that are numeric or string values, and when
// called returns true if actual < expected. Strings are compared lexicographically with '<'.
// Returns a matcher that checks if actual is greater than expected.
func LessThan[A Comparable](expected A) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("value less than <%v>", expected)
	matcher.Matches = func(actual A) bool {
		return actual < expected
	}
	return matcher
}
