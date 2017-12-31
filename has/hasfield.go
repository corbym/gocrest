package has

import (
	"reflect"
	"fmt"
	"github.com/corbym/gocrest"
)

// Naive implementation for testing if a struct has a particular field name. Does not check type.
// returns a matcher that will use reflect to check if the actual has the method given by expected
func FieldNamed(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("struct with function %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		typeOfActual := reflect.TypeOf(actual)
		matcher.Actual = fieldStringValue(typeOfActual)
		expectedName := reflect.ValueOf(expected).String()
		_, ok := typeOfActual.Elem().FieldByName(expectedName)
		return ok
	}
	return matcher
}

func fieldStringValue(actualType reflect.Type) string {
	description := actualType.Elem().Name() + "{"
	numFields := actualType.Elem().NumField()
	for x := 0; x < numFields; x++ {
		field := actualType.Elem().Field(x)
		description += fmt.Sprintf("%s %s", field.Name, field.Type.Name())
		if x != numFields-1 {
			description += " "
		}
	}
	description += "}"
	return description
}
