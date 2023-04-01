package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// Key is a matcher that checks if actual has a key == expected.
// Returns a matcher that matches when a map has key == expected
func Key[K comparable, V any](expected K) *gocrest.Matcher[map[K]V] {
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf("map has key '%v'", expected)
	matcher.Matches = func(actual map[K]V) bool {
		return hasKey(actual, expected)
	}
	return matcher
}

// AllKeys is a matcher that checks if map actual has all keys == expecteds.
// Returns a matcher that matches when a map has all keys == all expected.
func AllKeys[K comparable, V any](expected ...K) *gocrest.Matcher[map[K]V] {
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf("map has keys '%v'", expected)
	matcher.Matches = func(actual map[K]V) bool {
		for _, k := range expected {
			if !hasKey(actual, k) {
				return false
			}
		}
		return true
	}
	return matcher
}

func hasKey[K comparable, V any](actual map[K]V, expected K) bool {
	for k := range actual {
		if k == expected {
			return true
		}
	}
	return false
}
