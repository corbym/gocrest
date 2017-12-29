package is

import (
	"gocrest"
	"reflect"
)

// Matcher that matches if the actual is "empty"
// strings are empty if they are "", maps, arrays and slices are empty if len(actual) is 0
// pointers and interfaces are empty when nil
// Other types (int, float, bool) will cause the function to panic.
// Returns a matcher that evaluates true if actual is "empty"
func Empty() *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = "empty value"
	matcher.Matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		if actualValue, ok := actual.(string); ok {
			return actualValue == ""
		}
		if reflect.ValueOf(actual).Len() == 0 {
			return true
		}
		return false
	}
	return matcher
}
