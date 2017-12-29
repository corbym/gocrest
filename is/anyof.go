package is

import (
	"fmt"
	"gocrest"
)

//Takes some matchers and checks if at least one of the matchers return true
//returns a matcher that performs the the test on the input matchers
func AnyOf(allMatchers ... *gocrest.Matcher) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("any of (%s)", describe(allMatchers, "or"))
	matcher.Matches = func(actual interface{}) bool {
		for x := 0; x < len(allMatchers); x++ {
			if allMatchers[x].Matches(actual) {
				return true
			}
		}
		return false
	}
	return matcher
}
