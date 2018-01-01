package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// has.Key is a matcher that checks if actual has a key == expected.
// Panics when actual's Kind is not a map.
// returns a matcher that matches when a map has key == expected
func Key(expected interface{}) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("map has key '%s'", expected)
	matcher.Matches = func(actual interface{}) bool {
		return hasKey(actual, expected)
	}
	return matcher
}

// has.AllKeys is a matcher that checks if map actual has all keys == expecteds.
// Panics when actual's Kind is not a map.
// returns a matcher that matches when a map has all keys == all expected
func AllKeys(expected ...interface{}) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("map has keys '%s'", expected)
	matcher.Matches = func(actual interface{}) bool {
		keyValuesToMatch := reflect.ValueOf(correctExpectedValue(expected...))
		for x := 0; x < keyValuesToMatch.Len(); x++ {
			if !hasKey(actual, keyValuesToMatch.Index(x).Interface()) {
				return false
			}
		}
		return true
	}
	return matcher
}
func correctExpectedValue(expected ...interface{}) interface{} {
	kind := reflect.ValueOf(expected[0]).Kind()
	if kind == reflect.Slice {
		return expected[0]
	}
	return expected
}

func hasKey(actual interface{}, expected interface{}) bool {
	mapKeys := reflect.ValueOf(actual).MapKeys()
	for x := 0; x < len(mapKeys); x++ {
		if mapKeys[x].Interface() == expected {
			return true
		}
	}
	return false
}
