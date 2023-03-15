package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// FunctionNamed implementation for testing if a Type has a particular method name. Does not check parameters.
// Returns a matcher that will use reflect to check if the actual has the method given by expected.
func FunctionNamed(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("interface with function %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		typeOfActual := reflect.TypeOf(actual)
		matcher.Actual = actualStringValue(typeOfActual)
		expectedName := reflect.ValueOf(expected).String()
		_, ok := typeOfActual.Elem().MethodByName(expectedName)
		return ok
	}
	return matcher
}

func actualStringValue(actualType reflect.Type) string {
	description := actualType.Elem().Name() + "{"
	for x := 0; x < actualType.Elem().NumMethod(); x++ {
		description += actualType.Elem().Method(x).Name + "()"
	}
	description += "}"
	return description
}
