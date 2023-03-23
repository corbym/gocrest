package gocrest_test

import "fmt"

// StubTestingT stubs the testing.T interface for testing.
// It provides information on whether tests failed, and their output.
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
func (t *StubTestingT) Failed() bool {
	return t.failed
}
func (t *StubTestingT) Fail() {
	t.failed = true
}
func (t *StubTestingT) FailNow() {
	t.failed = true
}
func (t *StubTestingT) HasFailed() bool {
	return t.failed
}
func (t *StubTestingT) Helper() {
	//do nothing
}
