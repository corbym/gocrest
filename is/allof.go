package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

//AllOf takes some matchers and checks if all the matchers return true.
//Returns a matcher that performs the the test on the input matchers.
func AllOf(allMatchers ...*gocrest.Matcher) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Matches = func(actual interface{}) bool {
		matcher.Describe = fmt.Sprintf("all of (%s)", describe(allMatchers, "and"))
		for x := 0; x < len(allMatchers); x++ {
			if !allMatchers[x].Matches(actual) {
				return false
			}
		}
		return true
	}
	return matcher
}
