package gocrest

import "fmt"

func AllOf(allMatchers ... *Matcher) (*Matcher) {
	matcher := new(Matcher)
	matcher.matches = func(actual interface{}) bool {
		matcher.describe = fmt.Sprintf("all of (%s)", describe(allMatchers))
		for x := 0; x < len(allMatchers); x++ {

			if !allMatchers[x].matches(actual) {
				return false
			}
		}
		return true
	}
	return matcher
}
func describe(matchers []*Matcher) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += matchers[x].describe
		if x+1 < len(matchers) {
			description += " and "
		}
	}
	return description
}
