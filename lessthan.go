package gocrest

import (
	"reflect"
	"fmt"
)

// LessThan matcher compares two values that are numeric or string values, and when
// called returns true if actual < expected. Strings are compared lexicographically with '<'
// The matcher will always return false for unknown types.
// Actual and expected types must be the same underlying type, or the function will panic.
// Returns a matcher that checks if actual is greater than expected.
func LessThan(expected interface{}) *Matcher {
	matcher := new(Matcher)
	matcher.describe = fmt.Sprintf("value less than %v", expected)
	matcher.matches = func(actual interface{}) bool {
		actualValue := reflect.ValueOf(actual)
		expectedValue := reflect.ValueOf(expected)
		switch expected.(type) {
		case float32, float64:
			{
				return actualValue.Float() < expectedValue.Float()
			}
		case int, int8, int16, int32, int64:
			{
				return actualValue.Int() < expectedValue.Int()
			}
		case uint, uint8, uint16, uint32, uint64:
			{
				return actualValue.Uint() < expectedValue.Uint()
			}
		case string:
			return actualValue.String() < expectedValue.String()
		}
		return false
	}
	return matcher
}
