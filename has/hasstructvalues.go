package has

import (
	"fmt"
	"reflect"

	"github.com/corbym/gocrest"
)

type StructMatchers map[string]*gocrest.Matcher

func StructWithValues(expects map[string]*gocrest.Matcher) *gocrest.Matcher {
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

func valuesMatch(expects map[string]*gocrest.Matcher, actualValue reflect.Value) (string, string) {
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
