package is

import (
	"github.com/corbym/gocrest"
)

// Empty matches if the actual is "empty".
// 'string' values are empty if they are "", maps, arrays and slices are empty if len(actual) is 0.
// Pointers and interfaces are empty when nil.
// Other types (int, float, bool) will cause the function to panic.
// Returns a matcher that evaluates true if actual is "empty".
func Empty[K comparable, T any, A string | []T | map[K]T]() *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = "empty value"
	matcher.Matches = func(actual A) bool {
		return len(actual) == 0
	}
	return matcher
}
