package has

import (
	"github.com/corbym/gocrest"
	"reflect"
)

//TypeName returns true if the expected matches the actual Type's Name. Expected can be a matcher or a string.
// E.g. has.TypeName(EqualTo("pkg.Type)) would be true with instance of `type Type struct{}` in package name 'pkg'.
func TypeName(expected interface{}) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Matches = func(actual interface{}) bool {
		actualTypeName := reflect.TypeOf(actual).String()
		matcher.Actual = actualTypeName
		matcher.Describe = "has type "
		switch expected.(type) {
		case *gocrest.Matcher:
			m := expected.(*gocrest.Matcher)
			matches := m.Matches(actualTypeName)
			matcher.AppendActual(m.Actual)
			matcher.Describe += m.Describe
			return matches
		default:
			matcher.Describe += "<" + expected.(string) + ">"
			return actualTypeName == expected
		}
	}
	return matcher
}
