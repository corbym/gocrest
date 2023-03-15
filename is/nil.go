package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// Nil matches if the actual value is nil
func Nil() *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = "value that is <nil>"
	match.Matches = func(actual interface{}) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		if actual == nil {
			return true
		}
		switch reflect.TypeOf(actual).Kind() {
		case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
			return reflect.ValueOf(actual).IsNil()
		}
		return false // anything else is never nil (hopefully)
	}
	return match
}
