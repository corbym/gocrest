package is

import (
	"github.com/corbym/gocrest"
)

//Returns a matcher that returns logical not of the matcher given
func Not(matcher *gocrest.Matcher) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = "not(" + matcher.Describe + ")"
	match.Matches = func(actual interface{}) bool {
		return !matcher.Matches(actual)
	}
	return match
}
