package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// AllOf takes some matchers and checks if all the matchers return true.
// Returns a matcher that performs the test on the input matchers.
func AllOf(allMatchers ...*gocrest.Matcher) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("all of (%s)", describe("and", allMatchers))
	matcher.Matches = matchAll(allMatchers, matcher)
	return matcher
}

func matchAll(allMatchers []*gocrest.Matcher, allOf *gocrest.Matcher) func(actual interface{}) bool {
	return func(actual interface{}) bool {
		allOf.AppendActual(fmt.Sprintf("actual <%v>", actual))
		matches := true
		var failingMatchers []*gocrest.Matcher
		for x := 0; x < len(allMatchers); x++ {
			if !allMatchers[x].Matches(actual) {
				matches = false
				failingMatchers = append(failingMatchers, allMatchers[x])
			}
			allOf.AppendActual(allMatchers[x].Actual)
		}
		allOf.Describe = fmt.Sprintf("%s", describe("and", failingMatchers))
		return matches
	}
}
