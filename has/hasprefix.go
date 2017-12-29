package has

import (
	"gocrest"
	"strings"
	"reflect"
	"fmt"
)

//is.Prefix returns a matcher that matches if the given string is prefixed with the expected string
// panics if the actual is not a string
//returns a matcher that returns true if the above conditions are met
func Prefix(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("value with prefix %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual).String()
		return strings.HasPrefix(actualValue, expected)
	}
	return matcher
}
