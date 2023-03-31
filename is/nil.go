package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// Nil matches if the actual value is nil
func Nil[T any, A *T | any]() *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = "value that is <nil>"
	match.Matches = func(actual A) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		if any(actual) == nil {
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
