package has

import (
	"github.com/corbym/gocrest"
	"reflect"
)

func TypeName(expected interface{}) (matcher *gocrest.Matcher) {
	matcher = new(gocrest.Matcher)
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
		return false
	}
	return
}
