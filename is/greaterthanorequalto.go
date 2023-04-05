package is

import "github.com/corbym/gocrest"

// GreaterThanOrEqualTo is a shorthand matcher for anyOf(greaterThan(x), equalTo(x))
// Returns a matcher matching if actual >= expected (using deepEquals).
func GreaterThanOrEqualTo[A Comparable](expected A) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Matches = func(actual A) bool {
		anyOf := AnyOf(GreaterThan(expected), EqualTo(expected))
		anyOfMatches := anyOf.Matches(actual)
		matcher.Describe = anyOf.Describe
		return anyOfMatches
	}
	return matcher
}
