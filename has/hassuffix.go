package has

import (
	"gocrest"
	"strings"
	"reflect"
	"fmt"
)

//has.Suffix returns a matcher that matches if the given string is suffixed with the expected string
// panics if the actual is not a string
// uses strings.Suffix(act,exp) to evaluate strings
//returns a matcher that returns true if the above conditions are met
func Suffix(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("value with suffix %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual).String()
		return strings.HasSuffix(actualValue, expected)
	}
	return matcher
}
