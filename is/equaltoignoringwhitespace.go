package is

import (
	"github.com/corbym/gocrest"
	"strings"
)

// EqualToIgnoringWhitespace tests whether two strings have identical content without any whitespace
// comparison. For example:
//
// "a bc" is EqualToIgnoringWhitespace when compared with "a   b c"
// "a b c" is EqualToIgnoringWhitespace when compared with "a \nb \tc"
// "ab\tc" is EqualToIgnoringWhitespace when compared with "a \nb \tc"
// .. and so on.
func EqualToIgnoringWhitespace(expected string) (matcher *gocrest.Matcher[string]) {
	matcher = new(gocrest.Matcher[string])
	matcher.Matches = func(actual string) bool {
		expectedFields := strings.Join(strings.Fields(expected), "")
		actualFields := strings.Join(strings.Fields(actual), "")

		equalToMatcher := EqualTo(expectedFields)
		matcher.Describe = "ignoring whitespace value equal to <" + expected + ">"
		isEqualTo := equalToMatcher.Matches(actualFields)
		return isEqualTo
	}
	return matcher
}
