package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// AllOf takes some matchers and checks if all the matchers return true.
// Returns a matcher that performs the test on the input matchers.
func AllOf[A any](allMatchers ...*gocrest.Matcher[A]) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("all of (%s)", describe("and", allMatchers))
	matcher.Matches = matchAll(allMatchers, matcher)
	return matcher
}

func matchAll[A any](allMatchers []*gocrest.Matcher[A], allOf *gocrest.Matcher[A]) func(actual A) bool {
	return func(actual A) bool {
		allOf.AppendActual(fmt.Sprintf("actual <%v>", actual))
		matches := true
		var failingMatchers []*gocrest.Matcher[A]
		for _, m := range allMatchers {
			if !m.Matches(actual) {
				matches = false
				failingMatchers = append(failingMatchers, m)
			}
			allOf.AppendActual(m.Actual)
		}
		allOf.Describe = describe("and", failingMatchers)
		return matches
	}
}
