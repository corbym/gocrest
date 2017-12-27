package gocrest

import (
	"regexp"
	"reflect"
	"fmt"
)

func MatchesPattern(expected string) *Matcher {
	matcher := new(Matcher)
	matcher.describe = fmt.Sprintf("a value that matches pattern %s", expected)
	matcher.matches = func(actual interface{}) bool {
		compiledExp, err := regexp.Compile(expected)
		if err != nil {
			matcher.describe = err.Error()
			return false
		}
		return compiledExp.MatchString(reflect.ValueOf(actual).String())
	}
	return matcher
}
