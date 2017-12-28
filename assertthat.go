package gocrest

import "fmt"

//AssertThat calls a given matcher and fails the test with a message if the matcher doesn't match.
func AssertThat(t TestingT, actual interface{}, m *Matcher) {
	matches := m.matches(actual)
	if !matches {
		t.Errorf("expected: %s but was: %s", m.describe, actualAsString(m, actual))
	}
}

func actualAsString(matcher *Matcher, actual interface{}) string {
	if matcher.actual != "" {
		return matcher.actual
	}
	return fmt.Sprintf("%v", actual)
}
