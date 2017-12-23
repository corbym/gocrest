package gocrest

import (
	"reflect"
	"fmt"
)

func Contains(expected interface{}) *Matcher {
	match := new(Matcher)
	match.describe = fmt.Sprintf("something that contains %v", expected)
	match.matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		expectedValue := reflect.ValueOf(expected)
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
	return false
}

func
contains(value interface{}, list reflect.Value) bool {
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
	return len(contains) == actualValue.Len()
}
