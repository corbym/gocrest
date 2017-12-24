package gocrest

import (
	"fmt"
)

//Matcher to check if two values are equal.
//returns a matcher that will return true if two values are equal
func EqualTo(expected interface{}) *Matcher {
	match := new(Matcher)
	match.describe = fmt.Sprintf("value equal to %v", expected)
	match.matches = func(actual interface{}) bool {
		return expected == actual
	}

	return match
}
