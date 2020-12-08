package has

import (
	"fmt"
	"reflect"

	"github.com/corbym/gocrest"
)

// Type that can be passed to StructWithValues. Mapps Struct field names to a matcher
type StructMatchers map[string]*gocrest.Matcher

// Checks wether the actual struct matches all expectations passed as StructMatchers.
// This method can be used to check single struct fields in different ways or omit checking some struct fields at all.
// Panics if the actual value is not a struct.
// Panics if Structmatchers contains a key that can not be found in the actual struct.
func StructWithValues(expects StructMatchers) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Matches = func(actual interface{}) bool {

		actualValue := reflect.ValueOf(actual)
		switch actualValue.Kind() {
		case reflect.Struct:
			reason, actual := valuesMatch(expects, actualValue)

			match.Describe = reason
			match.Actual = actual
			return reason == "" && actual == ""

		default:
			panic("cannot determine type of variadic actual, " + actualValue.String())
		}
	}

	return match
}

func valuesMatch(expects StructMatchers, actualValue reflect.Value) (string, string) {
	for key, expect := range expects {
		v := actualValue.FieldByName(key)

		if !v.IsValid() {
			panic(fmt.Sprintf("Expect[%v] does not exist on actual struct", key))
		}

		result := expect.Matches(v.Interface())

		actual := expect.Actual

		if actual == "" {
			actual = v.String()
		}

		if !result {
			return fmt.Sprintf("expect[%v]: %v", key, expect.Describe), actual
		}
	}

	return "", ""
}
