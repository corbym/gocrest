package is

import (
	"fmt"
	"gocrest"
)

//Matcher to check if two values are equal.
//returns a matcher that will return true if two values are equal
func EqualTo(expected interface{}) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = fmt.Sprintf("value equal to %v", expected)
	match.Matches = func(actual interface{}) bool {
		return expected == actual
	}

	return match
}