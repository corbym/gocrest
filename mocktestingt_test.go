package gocrest

import "fmt"

type MockTestingT struct {
	failed         bool
	MockTestOutput string
}

func (t *MockTestingT) Logf(format string, args ...interface{}) {
	t.MockTestOutput = fmt.Sprintf(format, args...)
	t.failed = true
}

func (t *MockTestingT) Errorf(format string, args ...interface{}) {
	t.MockTestOutput = fmt.Sprintf(format, args...)
	t.failed = true
}

func (t *MockTestingT) FailNow() {
	t.failed = true
}
func (t *MockTestingT) HasFailed() bool {
	return t.failed
}
