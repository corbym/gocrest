package is

import (
	"fmt"
	"gocrest/base"
)

//Takes some matchers and checks if at least one of the matchers return true
//returns a matcher that performs the the test on the input matchers
func AnyOf(allMatchers ... *base.Matcher) *base.Matcher {
	matcher := new(base.Matcher)
	matcher.Matches = func(actual interface{}) bool {
		matcher.Describe = fmt.Sprintf("any of (%s)", describe(allMatchers, "or"))
		for x := 0; x < len(allMatchers); x++ {
			if allMatchers[x].Matches(actual) {
				return true
			}
		}
		return false
	}
	return matcher
}