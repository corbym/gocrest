package gocrest

import (
	"reflect"
)

func HasFunction(expected interface{}) *Matcher {
	matcher := new(Matcher)
	matcher.matches = func(actual interface{}) bool {
		_, ok := reflect.TypeOf(actual).MethodByName(reflect.ValueOf(expected).String())
		return ok
	}
	return matcher
}
