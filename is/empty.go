package is

import (
	"github.com/corbym/gocrest"
)

// Empty matches if the actual array or slice is len(actual) is 0.
// Returns a matcher that evaluates true if actual is "empty".
func Empty[A any]() *gocrest.Matcher[[]A] {
	matcher := new(gocrest.Matcher[[]A])
	matcher.Describe = "empty value"
	matcher.Matches = func(actual []A) bool {
		return len(actual) == 0
	}
	return matcher
}

// EmptyString matches if the actual string if they are "" (==len(0))
// Returns a matcher that evaluates true if actual is "empty".
func EmptyString() *gocrest.Matcher[string] {
	matcher := new(gocrest.Matcher[string])
	matcher.Describe = "empty value"
	matcher.Matches = func(actual string) bool {
		return len(actual) == 0
	}
	return matcher
}

// EmptyMap matches if the actual maps if len(actual) is 0.
// Returns a matcher that evaluates true if actual is "empty".
func EmptyMap[K comparable, V any]() *gocrest.Matcher[map[K]V] {
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = "empty value"
	matcher.Matches = func(actual map[K]V) bool {
		return len(actual) == 0
	}
	return matcher
}
