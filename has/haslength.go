package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// Length can be called with arrays and strings
// Returns a matcher that matches if the length matches the given criteria
func Length[V any, A []V | string](expected int) *gocrest.Matcher[A] {
	const description = "value with length %v"
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual A) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return lenOfActual == expected
	}
	return matcher
}

// MapLength can be called with maps
// Returns a matcher that matches if the length matches the given criteria
func MapLength[K comparable, V any](expected int) *gocrest.Matcher[map[K]V] {
	const description = "value with length %v"
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual map[K]V) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return lenOfActual == expected
	}
	return matcher
}

// LengthMatching can be called with arrays or strings
// Returns a matcher that matches if the length matches matcher passed in
func LengthMatching[V any, A []V | string](expected *gocrest.Matcher[int]) *gocrest.Matcher[A] {
	const description = "value with length %v"
	matcher := new(gocrest.Matcher[A])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual A) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return expected.Matches(lenOfActual)

	}
	return matcher
}

// MapLengthMatching can be called with maps
// Returns a matcher that matches if the length matches the given matcher
func MapLengthMatching[K comparable, V any](expected *gocrest.Matcher[int]) *gocrest.Matcher[map[K]V] {
	const description = "value with length %v"
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual map[K]V) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return expected.Matches(lenOfActual)

	}
	return matcher
}
