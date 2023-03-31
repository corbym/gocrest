package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

func StringContaining(expected ...string) *gocrest.Matcher[string] {
	match := new(gocrest.Matcher[string])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual string) bool {
		for i := 0; i < len(expected); i++ {
			if !strings.Contains(actual, expected[i]) {
				return false
			}
		}
		return true
	}
	return match
}

// MapContaining finds if x is contained as value in y.
// Acts like "ContainsAll", all elements given must be present (or must match) in actual in the same order as the expected values.
func MapContaining[K comparable, V comparable, A map[K]V](expected A) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual A) bool {
		return mapActualContainsExpected(expected, actual)
	}
	return match
}
func MapContainingValues[K comparable, V comparable, A map[K]V](expected ...V) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual A) bool {
		return mapActualContainsExpectedValues(expected, actual)
	}
	return match
}
func MapMatchingValues[K comparable, V comparable, A map[K]V](expected ...*gocrest.Matcher[V]) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = descriptionForMatchers(expected)
	match.Matches = func(actual A) bool {
		return mapActualMatchesExpected(expected, actual)
	}
	return match
}

func descriptionForMatchers[A any](expected []*gocrest.Matcher[A]) string {
	var description = ""
	for x := 0; x < len(expected); x++ {
		description += expected[x].Describe
		if x < len(expected)-1 {
			description += " and "
		}
	}
	return description
}

// ArrayContaining finds if x is contained in y.
// Acts like "ContainsAll", all elements given must be present (or must match) in actual in the same order as the expected values.
func ArrayContaining[T comparable, A []T](expected ...T) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual A) bool {
		return listContains(expected, actual)
	}
	return match
}
func ArrayMatching[T comparable, A []T](expected ...*gocrest.Matcher[T]) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual A) bool {
		return listMatches(expected, actual)
	}
	return match
}
func mapActualContainsExpected[K comparable, V comparable](expected map[K]V, actual map[K]V) bool {
	expectedKeys := make([]K, 0, len(expected))
	for k := range expected {
		expectedKeys = append(expectedKeys, k)
	}
	contains := make(map[V]bool)
	for i := 0; i < len(expectedKeys); i++ {
		val := actual[expectedKeys[i]]
		if val == expected[expectedKeys[i]] {
			contains[val] = true
		}
	}
	return len(contains) == len(expectedKeys)
}
func mapActualContainsExpectedValues[K comparable, V comparable](expected []V, actual map[K]V) bool {
	contains := make(map[V]bool)
	for i := 0; i < len(expected); i++ {
		for k, v := range actual {
			if actual[k] == expected[i] {
				contains[v] = true
				break
			}
		}
	}
	return len(contains) == len(expected)
}
func mapActualMatchesExpected[K comparable, V comparable, A map[K]V](expected []*gocrest.Matcher[V], actual A) bool {
	actualKeys := make([]K, 0, len(actual))
	for k := range actual {
		actualKeys = append(actualKeys, k)
	}
	contains := make(map[interface{}]bool)
	for i := 0; i < len(expected); i++ {
		for _, v := range actual {
			if expected[i].Matches(v) {
				contains[v] = true
			}
		}
	}
	return len(contains) == len(expected)
}

func listContains[T comparable, A []T](expected A, actualValue A) bool {
	contains := make(map[interface{}]bool)
	for i := 0; i < len(expected); i++ {
		for y := 0; y < len(actualValue); y++ {
			exp := expected[i]
			act := actualValue[y]
			if exp == act {
				contains[act] = true
			}
		}
	}
	return len(contains) == len(expected)
}
func listMatches[T comparable](expected []*gocrest.Matcher[T], actualValue []T) bool {
	contains := make(map[interface{}]bool)
	for i := 0; i < len(expected); i++ {
		for y := 0; y < len(actualValue); y++ {
			exp := expected[i]
			act := actualValue[y]
			if exp.Matches(act) {
				contains[act] = true
			}
		}
	}
	return len(contains) == len(expected)
}

func descriptionFor[T any, A []T](expected A) string {
	var description = ""
	for x := 0; x < len(expected); x++ {
		description += fmt.Sprintf("<%s>", expected[x])
		if x < len(expected)-1 {
			description += " and "
		}
	}
	return description
}
