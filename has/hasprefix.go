package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

// Prefix returns a matcher that matches if the given string is prefixed with the expected string
// Uses strings.Prefix(act, exp) to evaluate strings.
// Returns a matcher that returns true if the above conditions are met
func Prefix(expected string) *gocrest.Matcher[string] {
	matcher := new(gocrest.Matcher[string])
	matcher.Describe = fmt.Sprintf("value with prefix %s", expected)
	matcher.Matches = func(actual string) bool {
		return strings.HasPrefix(actual, expected)
	}
	return matcher
}
