package then

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
	"sync"
	"time"
)

// FailureLog is a type that holds temporary test data whilst Eventually runs
type FailureLog struct {
	TestOutput string
	failed     bool
}

// RecordingTestingT is a testingT interface that holds temporary test data whilst Eventually runs. Only used with Eventually
type RecordingTestingT struct {
	sync.Mutex
	gocrest.TestingT
	failures []FailureLog
}

// Errorf appends an error message into the failure buffer. Only used with Eventually
func (t *RecordingTestingT) Errorf(format string, args ...interface{}) {
	t.Lock()
	defer t.Unlock()
	t.failures = append(t.failures, FailureLog{
		fmt.Sprintf(format, args...),
		true,
	})
}

// Fail calls Errorf with a generic failure message in case the test is failed from inside the value under test.
// Only used with Eventually
func (t *RecordingTestingT) Fail() {
	t.Errorf("Unknown call to Fail")
}

// FailNow calls Errorf with a generic failure message in case the test is failed from inside the value under test.
// Only used with Eventually
func (t *RecordingTestingT) FailNow() {
	t.Errorf("Unknown call to FailNow")
}

// FailedTestOutputs retrieves the list of recorded testing.T failures from the assertions passed in for Eventually evaluation.
// Only used with Eventually
func (t *RecordingTestingT) FailedTestOutputs() []string {
	t.Lock()
	defer t.Unlock()
	var logs []string
	for _, failure := range t.failures {
		logs = append(logs, failure.TestOutput)
	}
	return logs
}

// Failing determines whether any of the Eventually assertions are still failing.
// Only used with Eventually
func (t *RecordingTestingT) Failing() bool {
	t.Lock()
	defer t.Unlock()
	for _, failure := range t.failures {
		if failure.failed {
			return true
		}
	}
	return false
}

// Latest is used internally by Eventually to record the last recorded test output from assertions passed into it
// Only used with Eventually
type Latest struct {
	sync.Mutex
	latestValue RecordingTestingT
}

// Get is used internally by Eventually to Get the last recorded test output from assertions passed into it
// Only used with Eventually
func (l *Latest) Get() RecordingTestingT {
	l.Lock()
	defer l.Unlock()
	return l.latestValue
}

// Merge is used internally by Eventually to Merge the latest recorded test output from assertions passed into it with the last one
// Only used with Eventually
func (l *Latest) Merge(updated RecordingTestingT) RecordingTestingT {
	l.Lock()
	defer l.Unlock()
	var mergedFailures []FailureLog
	for i, failure := range l.latestValue.failures {
		if failure.failed {
			if i < len(updated.failures) {
				mergedFailures = append(mergedFailures, updated.failures[i])
			}
		} else {
			mergedFailures = append(mergedFailures, failure)
		}
	}
	if l.latestValue.failures == nil {
		mergedFailures = updated.failures
	}
	merged := RecordingTestingT{
		failures: mergedFailures,
		TestingT: l.latestValue.TestingT,
	}
	return merged
}

// Eventually retries a set of assertions passed into it from a function.
// waitFor is the amount of time to wait for the assertions to pass.
// tick is the amount of time between running the assertions
// assertions are the function containing the assertions.
// e.g:
// ```
//
//	then.Eventually(t, time.Second*5, time.Second, func(eventually gocrest.TestingT) {
//		then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
//	})
//
// ```
func Eventually(t gocrest.TestingT, waitFor, tick time.Duration, assertions func(eventually gocrest.TestingT)) {

	t.Helper()
	channel := make(chan RecordingTestingT, 1)
	defer close(channel)

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	var latestValue = new(Latest)
	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			latestRecordingT := latestValue.Get()
			t.Errorf(fmt.Sprintf("Eventually Failed after %s: \n", waitFor) + strings.Join(latestRecordingT.FailedTestOutputs(), "\n"))
			return
		case <-tick:
			tick = nil
			go func() {
				recordedTesting := RecordingTestingT{
					TestingT: t,
					failures: []FailureLog{},
				}
				assertions(&recordedTesting)
				channel <- recordedTesting
			}()
		case value := <-channel:
			if !value.Failing() {
				return
			}
			latestValue.latestValue = latestValue.Merge(value)
			tick = ticker.C
		}
	}

}

// WithinTenSeconds is a shortcut for a ten-second Eventually call with one second tick
func WithinTenSeconds(t gocrest.TestingT, assertions func(eventually gocrest.TestingT)) {
	Eventually(t, time.Duration(10)*time.Second, time.Duration(1)*time.Second, assertions)
}

// WithinFiveSeconds is a shortcut for a five-second Eventually call with one second tick
func WithinFiveSeconds(t gocrest.TestingT, assertions func(eventually gocrest.TestingT)) {
	Eventually(t, time.Duration(10)*time.Second, time.Duration(1)*time.Second, assertions)
}
