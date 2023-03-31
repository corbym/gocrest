package then

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// AssertThat calls a given matcher and fails the test with a message if the matcher doesn't match.
func AssertThat[A any](t gocrest.TestingT, actual A, m *gocrest.Matcher[A]) {
	t.Helper()
	matches := m.Matches(actual)
	if !matches {
		t.Errorf("%s\nExpected: %s"+
			"\n     but: <%s>\n",
			m.ReasonString,
			m.Describe,
			actualAsString(m, actual),
		)
	}
}

func actualAsString[A any](matcher *gocrest.Matcher[A], actual A) string {
	if matcher.Actual != "" {
		return matcher.Actual
	}
	return fmt.Sprintf("%v", actual)
}
