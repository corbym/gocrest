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
	matcher.Matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		actualValue := reflect.ValueOf(actual)
		if actualValue.Kind() == reflect.String {
			return actualValue.String() == ""
		}
		if actualValue.Len() == 0 {
			return true
		}
		return false
	}
	return matcher
}
