package is

import "gocrest"

// Short hand matcher for anyOf(greaterThan(x), equalTo(x))
// returns a matcher matching if actual >= expected (using deepEquals)
func GreaterThanOrEqualTo(expected interface{}) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Matches = func(actual interface{}) bool {
		anyOf := AnyOf(GreaterThan(expected), EqualTo(expected))
		anyOfMatches := anyOf.Matches(actual)
		matcher.Describe = anyOf.Describe
		return anyOfMatches
	}
	return matcher
}
