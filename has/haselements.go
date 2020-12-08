package has

import (
	"fmt"
	"reflect"

	"github.com/corbym/gocrest"
)

// Checks wether each element of the array/slice matches each exptectation passed to ElementsWith
// Panics if the actual is not an array/slice
func ElementsWith(expects ...*gocrest.Matcher) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Matches = func(actual interface{}) bool {

		actualValue := reflect.ValueOf(actual)
		switch actualValue.Kind() {
		case reflect.Array, reflect.Slice:
			reason, actual := elementsMatch(expects, actualValue)

			match.Describe = reason
			match.Actual = actual
			return reason == "" && actual == ""

		default:
			panic("cannot determine type of variadic actual, " + actualValue.String())
		}
	}

	return match
}

func elementsMatch(expects []*gocrest.Matcher, actualValue reflect.Value) (string, string) {
	for i := 0; i < actualValue.Len(); i++ {
		for j, expect := range expects {
			result := expect.Matches(actualValue.Index(i).Interface())

			actual := expect.Actual

			if actual == "" {
				actual = actualValue.Index(i).String()
			}

			if !result {
				return fmt.Sprintf("expect[%v]: %v", j, expect.Describe), fmt.Sprintf("actual[%v]: <%v>", i, actual)
			}
		}
	}

	return "", ""
}
