package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"regexp"
)

// Matches if actual string matches the expected regex
// string provided must be a valid for compilation with regexp.Compile
// returns a matcher that uses the expected for a regex to match the actual value
func MatchForPattern(expected string) *gocrest.Matcher {
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf("a value that matches pattern %s", expected)
	matcher.Matches = func(actual interface{}) bool {
		compiledExp, err := regexp.Compile(expected)
		if err != nil {
			matcher.Describe = err.Error()
			return false
		}
		return compiledExp.MatchString(actual.(string))
	}
	return matcher
}
