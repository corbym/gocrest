package then

import (
	"fmt"
	"gocrest"
)

//AssertThat calls a given matcher and fails the test with a message if the matcher doesn't match.
func AssertThat(t gocrest.TestingT, actual interface{}, m *gocrest.Matcher) {
	matches := m.Matches(actual)
	if !matches {
		t.Errorf("expected: %s but was: %s", m.Describe, actualAsString(m, actual))
	}
}

func actualAsString(matcher *gocrest.Matcher, actual interface{}) string {
	if matcher.Actual != "" {
		return matcher.Actual
	}
	return fmt.Sprintf("%v", actual)
}