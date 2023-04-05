package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

const description = "value with length %v"

// Length can be called with arrays
// Returns a matcher that matches if the length matches the given criteria
func Length[T any](expected int) *gocrest.Matcher[[]T] {
	matcher := new(gocrest.Matcher[[]T])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual []T) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return lenOfActual == expected
	}
	return matcher
}

// StringLength can be called with strings
// Returns a matcher that matches if the length matches the given criteria
func StringLength(expected int) *gocrest.Matcher[string] {
	matcher := new(gocrest.Matcher[string])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual string) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return lenOfActual == expected
	}
	return matcher
}

// MapLength can be called with maps
// Returns a matcher that matches if the length matches the given criteria
func MapLength[K comparable, V any](expected int) *gocrest.Matcher[map[K]V] {
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual map[K]V) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return lenOfActual == expected
	}
	return matcher
}

// LengthMatching can be called with arrays
// Returns a matcher that matches if the length matches matcher passed in
func LengthMatching[A any](expected *gocrest.Matcher[int]) *gocrest.Matcher[[]A] {
	matcher := new(gocrest.Matcher[[]A])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual []A) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return expected.Matches(lenOfActual)

	}
	return matcher
}

// StringLengthMatching can be called with arrays or strings
// Returns a matcher that matches if the length matches matcher passed in
func StringLengthMatching(expected *gocrest.Matcher[int]) *gocrest.Matcher[string] {
	matcher := new(gocrest.Matcher[string])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual string) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return expected.Matches(lenOfActual)

	}
	return matcher
}

// MapLengthMatching can be called with maps
// Returns a matcher that matches if the length matches the given matcher
func MapLengthMatching[K comparable, V any](expected *gocrest.Matcher[int]) *gocrest.Matcher[map[K]V] {
	matcher := new(gocrest.Matcher[map[K]V])
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual map[K]V) bool {
		lenOfActual := len(actual)
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		return expected.Matches(lenOfActual)

	}
	return matcher
}
