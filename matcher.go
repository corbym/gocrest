package gocrest

type Matcher struct {
	matches  func(actual interface{}) bool
	describe string
}

type TestingT interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
}

//AssertThat calls a given matcher and fails with a message if the matcher doesn't match.
func AssertThat(t TestingT, actual interface{}, m *Matcher) {
	matches := m.matches(actual)
	if !matches {
		t.Errorf("expected: %s but was: %v", m.describe, actual)
	}
}
