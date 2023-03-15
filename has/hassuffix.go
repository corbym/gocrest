package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

// Suffix returns a matcher that matches if the given string is suffixed with the expected string.
// Panics if the actual is not a string.
// Uses strings.HasSuffix(act,exp) to evaluate strings.
// Returns a matcher that returns true if the above conditions are met.
func Suffix(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("value with suffix %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		return strings.HasSuffix(actual.(string), expected)
	}
	return matcher
}
