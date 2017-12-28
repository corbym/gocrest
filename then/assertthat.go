package then

import (
	"fmt"
	"gocrest/base"
)

//AssertThat calls a given matcher and fails the test with a message if the matcher doesn't match.
func AssertThat(t base.TestingT, actual interface{}, m *base.Matcher) {
	matches := m.Matches(actual)
	if !matches {
		t.Errorf("expected: %s but was: %s", m.Describe, actualAsString(m, actual))
	}
}

func actualAsString(matcher *base.Matcher, actual interface{}) string {
	if matcher.Actual != "" {
		return matcher.Actual
	}
	return fmt.Sprintf("%v", actual)
}