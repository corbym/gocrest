package gocrest

import (
	"reflect"
	"fmt"
)

func HasFunctionNamed(expected string) *Matcher {
	matcher := new(Matcher)
	matcher.describe = fmt.Sprintf("interface with function %s", expected)
	matcher.matches = func(actual interface{}) bool {
		typeOfActual := reflect.TypeOf(actual)
		matcher.resolvedActual = actualStringValue(typeOfActual)
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
