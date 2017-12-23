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

func AssertThat(t TestingT, actual interface{}, m *Matcher) bool {
	matches := m.matches(actual)
	if ! matches {
		t.Errorf("Expected: %s\n Actual: %v", m.describe, actual)
	}
	return matches
}
