package is

import (
	"reflect"
	"fmt"
	"strings"
	"gocrest"
)

// ValueContaining finds if x is contained in y.
// Acts like "ContainsAll", all elements given must be present in actual
// If "expected" is an array or slice, we assume that actual is the same type
// For maps, the expected must also be a map and contains is true if both maps contain all key,values in expected.
// For string, behaves like strings.Contains
// Will panic if types cannot be converted correctly.
// returns the Matcher that returns true if found
func ValueContaining(expected interface{}) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = fmt.Sprintf("something that contains %v", expected)
	match.Matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		expectedValue := reflect.ValueOf(expected)

		if expectedValue.Kind() == reflect.String && actualValue.Kind() == reflect.String {
			return strings.Contains(actualValue.String(), expectedValue.String())
		}

		switch expectedValue.Kind() {
		case reflect.Array, reflect.Slice:
			{
				return listContains(expectedValue, actualValue)
			}
		case reflect.Map:
			{
				return mapContains(expectedValue, actualValue)
			}
		default:
			{
				return contains(expected, actualValue)
			}
		}
	}
	return match
}

func mapContains(expected reflect.Value, actual reflect.Value) bool {
	keys := expected.MapKeys()
	contains := make(map[interface{}]bool)
	for i := 0; i < len(keys); i++ {
		val := actual.MapIndex(keys[i])
		if val.IsValid() {
			if val.Interface() == expected.MapIndex(keys[i]).Interface() {
				contains[val] = true
			}
		}
	}
	return len(contains) == len(expected.MapKeys())
}

func contains(value interface{}, list reflect.Value) bool {
	for i := 0; i < list.Len(); i++ {
		if list.Index(i).Interface() == value {
			return true
		}
	}
	return false
}

func listContains(expectedValue reflect.Value, actualValue reflect.Value) bool {
	contains := make(map[interface{}]bool)
	for i := 0; i < expectedValue.Len(); i++ {
		for y := 0; y < actualValue.Len(); y++ {
			exp := expectedValue.Index(i).Interface()
			act := actualValue.Index(y).Interface()
			if exp == act {
				contains[act] = true
			}
		}
	}
	return len(contains) == expectedValue.Len()
}