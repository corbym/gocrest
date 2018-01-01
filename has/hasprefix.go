package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

//Prefix returns a matcher that matches if the given string is prefixed with the expected string
//Function panics if the actual is not a string.
//Uses strings.Prefix(act, exp) to evaluate strings.
//Returns a matcher that returns true if the above conditions are met
func Prefix(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("value with prefix %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		return strings.HasPrefix(actual.(string), expected)
	}
	return matcher
}
