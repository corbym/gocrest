package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// StructMatchers Type that can be passed to StructWithValues. Map Struct field names to a matcher
type StructMatchers[A any] map[string]*gocrest.Matcher[A]

// StructWithValues Checks whether the actual struct matches all expectations passed as StructMatchers.
// This method can be used to check single struct fields in different ways or omit checking some struct fields at all.
// Will automatically de-reference pointers.
// Panics if the actual value is not a struct.
// Panics if StructMatchers contains a key that can not be found in the actual struct.
// Panics if StructMatchers contains a key that is unexported.
func StructWithValues[A any, B any](expects StructMatchers[B]) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("struct values to match {%s}", describeStructMatchers(expects))

	for _, e := range expects {
		match.AppendActual(e.Actual)
	}

	match.Matches = func(actual A) bool {

		actualValue := reflect.ValueOf(actual)
		if actualValue.Kind() == reflect.Ptr {
			actualValue = actualValue.Elem()
		}
		switch actualValue.Kind() {
		case reflect.Struct:
			for key, expect := range expects {
				v := actualValue.FieldByName(key)

				if !v.IsValid() {
					panic(fmt.Sprintf("Expect[%v] does not exist on actual struct", key))
				}
				var matches = false
				if value, ok := v.Interface().(B); ok {
					matches = expect.Matches(value)
				}
				if !matches {
					return false
				}
			}
			return true
		}
		panic("cannot determine type of actual, " + actualValue.String())
	}

	return match
}

func describeStructMatchers[A any](matchers StructMatchers[A]) string {
	description := ""

	bindCount := 0

	for key, matcher := range matchers {
		description += fmt.Sprintf("\"%v\": %v", key, matcher.Describe)

		if bindCount < len(matchers)-1 {
			description += " and "
		}

		bindCount++
	}

	return description
}
