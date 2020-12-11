package has

import (
	"fmt"
	"reflect"

	"github.com/corbym/gocrest"
)

// Checks whether the nth element of the array/slice matches the nth expectation passed
// Panics if the actual is not an array/slice
// Panics if the count of the expectations does not match the array's/slice's length
func EveryElement(expects ...*gocrest.Matcher) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = fmt.Sprintf("elements to match %s", describe(expects, "and"))

	for _, e := range expects {
		match.AppendActual(e.Actual)
	}

	match.Matches = func(actual interface{}) bool {

		actualValue := reflect.ValueOf(actual)
		switch actualValue.Kind() {
		case reflect.Array, reflect.Slice:

			if actualValue.Len() != len(expects) {
				panic(fmt.Sprintf("cannot match expectations (length %v) to actuals (length %v)", len(expects), actualValue.Len()))
			}

			for i := 0; i < actualValue.Len(); i++ {
				result := expects[i].Matches(actualValue.Index(i).Interface())

				if !result {
					return false
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
		description += fmt.Sprintf("[%v]:%v", x, matchers[x].Describe)
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
