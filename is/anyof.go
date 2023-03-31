package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// AnyOf takes some matchers and checks if at least one of the matchers return true.
// Returns a matcher that performs the test on the input matchers.
func AnyOf[A any](allMatchers ...*gocrest.Matcher[A]) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("any of (%s)", describe("or", allMatchers))
	matcher.Matches = anyMatcherMatches(allMatchers, matcher)
	return matcher
}

func anyMatcherMatches[A any](allMatchers []*gocrest.Matcher[A], anyOf *gocrest.Matcher[A]) func(actual A) bool {
	return func(actual A) bool {
		matches := false
		anyOf.AppendActual(fmt.Sprintf("actual <%v>", actual))
		for x := 0; x < len(allMatchers); x++ {
			if allMatchers[x].Matches(actual) {
				matches = true
			}
			anyOf.AppendActual(allMatchers[x].Actual)
		}
		return matches
	}
}
