package gocrest

// TestingT supplies a convenience interface that matches the testing.T interface.
type TestingT interface {
	Logf(format string, args ...any)
	Errorf(format string, args ...any)
	Failed() bool
	Fail()
	FailNow()
	Helper()
}
