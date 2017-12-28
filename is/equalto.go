package is

import (
	"fmt"
	"gocrest/base"
)

//Matcher to check if two values are equal.
//returns a matcher that will return true if two values are equal
func EqualTo(expected interface{}) *base.Matcher {
	match := new(base.Matcher)
	match.Describe = fmt.Sprintf("value equal to %v", expected)
	match.Matches = func(actual interface{}) bool {
		return expected == actual
	}

	return match
}