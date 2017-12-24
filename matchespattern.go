package gocrest

import (
	"regexp"
	"reflect"
	"fmt"
)

func MatchesPattern(expected interface{}) (*Matcher) {
	matcher := new(Matcher)
	matcher.describe = fmt.Sprintf("a value that matches pattern %v", expected)
	matcher.matches = func(actual interface{}) bool {
		compiledExp, err := regexp.Compile(reflect.ValueOf(expected).String())
		if err != nil {
			matcher.describe = err.Error()
			return false
		}
		return compiledExp.MatchString(reflect.ValueOf(actual).String())
	}
	return matcher
}
