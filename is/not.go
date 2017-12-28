package is

import "gocrest/base"

//Returns a matcher that returns logical not of the matcher given
func Not(matcher *base.Matcher) *base.Matcher {
	match := new(base.Matcher)
	match.Describe = "not(" + matcher.Describe + ")"
	match.Matches = func(actual interface{}) bool {
		return !matcher.Matches(actual)
	}
	return match
}
