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
	match.Describe = fmt.Sprintf("elements to match all of (%s)", describe(expects, "and"))

	for _, e := range expects {
		match.AppendActual(e.Actual)
	}

	match.Matches = func(actual interface{}) bool {

		actualValue := reflect.ValueOf(actual)
		switch actualValue.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < actualValue.Len(); i++ {
				for _, expect := range expects {
					result := expect.Matches(actualValue.Index(i).Interface())

					if !result {
						return false
					}
				}
			}

			return true

		default:
			panic("cannot determine type of variadic actual, " + actualValue.String())
		}
	}

	return match
}

func describe(matchers []*gocrest.Matcher, conjunction string) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += matchers[x].Describe
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
