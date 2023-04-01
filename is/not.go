package is

import (
	"github.com/corbym/gocrest"
)

// Not negates the given matcher.
// Returns a matcher that returns logical not of the matcher given.
func Not[A any](matcher *gocrest.Matcher[A]) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = "not(" + matcher.Describe + ")"
	match.Matches = func(actual A) bool {
		matches := !matcher.Matches(actual)
		match.Actual = matcher.Actual
		return matches
	}
	return match
}
