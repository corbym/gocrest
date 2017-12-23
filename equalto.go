package gocrest

import (
	"fmt"
)

func EqualTo(expected interface{}) *Matcher {
	match := new(Matcher)
	match.describe = fmt.Sprintf("value equal to %v", expected)
	match.matches = func(actual interface{}) bool {
		return expected == actual
	}

	return match
}
