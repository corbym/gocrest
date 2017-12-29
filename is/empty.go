package is

import (
	"gocrest"
	"reflect"
)

func Empty() *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Matches = func(actual interface{}) bool {
		if actual == nil {
			return true
		}
		actualValue := reflect.ValueOf(actual)
		if actualValue.Kind() == reflect.String {
			return actualValue.String() == ""
		}
		if actualValue.Len() == 0 {
			return true
		}
		return false
	}
	return matcher
}
