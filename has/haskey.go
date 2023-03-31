package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// Key is a matcher that checks if actual has a key == expected.
// Panics when actual's Kind is not a map.
// Returns a matcher that matches when a map has key == expected
func Key[K comparable, V any, A map[K]V](expected K) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("map has key '%s'", expected)
	matcher.Matches = func(actual A) bool {
		return hasKey(actual, expected)
	}
	return matcher
}

// AllKeys is a matcher that checks if map actual has all keys == expecteds.
// Panics when actual's Kind is not a map.
// Returns a matcher that matches when a map has all keys == all expected.
func AllKeys[K comparable, V any, A map[K]V](expected ...K) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf("map has keys '%s'", expected)
	matcher.Matches = func(actual A) bool {
		for _, k := range expected {
			if !hasKey(actual, k) {
				return false
			}
		}
		return true
	}
	return matcher
}

func hasKey[K comparable, V any, A map[K]V](actual A, expected K) bool {
	for k, _ := range actual {
		if k == expected {
			return true
		}
	}
	return false
}
