package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
)

// StringContaining finds if all x's are contained as value in y.
// Acts like "ContainsAll", all elements given must be present.
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

// MapContaining finds if all map[k] 's value V is contained as a value of actual[k]
// Acts like "ContainsAll", all elements given must be present in actual in the same order as the expected values.
func MapContaining[K comparable, V comparable](expected map[K]V) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual map[K]V) bool {
		return mapActualContainsExpected(expected, actual)
	}
	return match
}

// MapContainingValues finds if all values V is contained as a value of actual[k]
// Acts like "ContainsAll", all elements given must be present in actual in the same order as the expected values.
func MapContainingValues[K comparable, V comparable](expected ...V) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual map[K]V) bool {
		return mapActualContainsExpectedValues(expected, actual)
	}
	return match
}

// MapMatchingValues finds if all values V is match a value of actual[k]
// Acts like "ContainsAll", all elements given must match in actual in the same order as the expected values.
func MapMatchingValues[K comparable, V comparable](expected ...*gocrest.Matcher[V]) *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = descriptionForMatchers(expected)
	match.Matches = func(actual map[K]V) bool {
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

// ArrayContaining finds if all x's are contained in y.
// Acts like "ContainsAll", all elements given must be present in actual.
func ArrayContaining[A comparable](expected ...A) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual []A) bool {
		return listContains(expected, actual)
	}
	return match
}

// ArrayMatching finds if all x's are matched in y.
// Acts like "ContainsAll", all elements given must be present in actual.
func ArrayMatching[A comparable](expected ...*gocrest.Matcher[A]) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("something that contains %v", descriptionFor(expected))
	match.Matches = func(actual []A) bool {
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
	contains := make(map[T]bool)
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
		description += fmt.Sprintf("<%v>", expected[x])
		if x < len(expected)-1 {
			description += " and "
		}
	}
	return description
}
