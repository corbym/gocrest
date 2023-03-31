package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// Length can be called with arrays, maps, *gocrest.Matcher and strings but not numeric types.
// has.Length(is.GreaterThan(x)) is a valid call.
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
func MapLength[K comparable, V any, A map[K]V](expected int) *gocrest.Matcher[A] {
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
func LengthMatching[A []any | string](expected *gocrest.Matcher[int]) *gocrest.Matcher[A] {
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
func MapLengthMatching[K comparable, V any, A map[K]V](expected *gocrest.Matcher[int]) *gocrest.Matcher[A] {
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
