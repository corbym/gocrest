package gocrest

import (
	"reflect"
	"fmt"
)

func GreaterThan(expected interface{}) *Matcher {
	matcher := new(Matcher)
	matcher.describe = fmt.Sprintf("value greater than %v", expected)
	matcher.matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		expectedValue := reflect.ValueOf(expected)
		switch expected.(type) {
		case float32, float64:
			{
				return actualValue.Float() > expectedValue.Float()
			}
		case int, int8, int16, int32, int64:
			{
				return actualValue.Int() > expectedValue.Int()
			}
		case uint, uint8, uint16, uint32, uint64:
			{
				return actualValue.Uint() > expectedValue.Uint()
			}
		}
		return false
	}

	return matcher
}
