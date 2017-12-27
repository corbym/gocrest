package gocrest

import "fmt"

type StubTestingT struct {
	failed         bool
	MockTestOutput string
}

func (t *StubTestingT) Logf(format string, args ...interface{}) {
	t.MockTestOutput = fmt.Sprintf(format, args...)
	t.failed = true
}

func (t *StubTestingT) Errorf(format string, args ...interface{}) {
	t.MockTestOutput = fmt.Sprintf(format, args...)
	t.failed = true
}

func (t *StubTestingT) FailNow() {
	t.failed = true
}
func (t *StubTestingT) HasFailed() bool {
	return t.failed
}
