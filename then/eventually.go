package then

import (
	"fmt"
	"github.com/corbym/gocrest"
	"strings"
	"sync"
	"time"
)

type FailureLog struct {
	TestOutput string
	failed     bool
}
type RecordingTestingT struct {
	sync.Mutex
	gocrest.TestingT
	failures []FailureLog
}

func (t *RecordingTestingT) Errorf(format string, args ...interface{}) {
	t.Lock()
	defer t.Unlock()
	t.failures = append(t.failures, FailureLog{
		fmt.Sprintf(format, args...),
		true,
	})
}
func (t *RecordingTestingT) Fail() {
	t.Errorf("Unknown call to Fail")
}
func (t *RecordingTestingT) FailNow() {
	t.Errorf("Unknown call to FailNow")
}
func (t *RecordingTestingT) FailedTestOutputs() []string {
	t.Lock()
	defer t.Unlock()
	var logs []string
	for _, failure := range t.failures {
		logs = append(logs, failure.TestOutput)
	}
	return logs
}
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

type Latest struct {
	sync.Mutex
	latestValue RecordingTestingT
}

func (l *Latest) Get() RecordingTestingT {
	l.Lock()
	defer l.Unlock()
	return l.latestValue
}
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
func Eventually(t gocrest.TestingT, waitFor time.Duration, tick time.Duration, assertions func(eventually gocrest.TestingT)) {

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
func WithinTenSeconds(t gocrest.TestingT, assertions func(eventually gocrest.TestingT)) {
	Eventually(t, time.Duration(10)*time.Second, time.Duration(1)*time.Second, assertions)
}
func WithinFiveSeconds(t gocrest.TestingT, assertions func(eventually gocrest.TestingT)) {
	Eventually(t, time.Duration(10)*time.Second, time.Duration(1)*time.Second, assertions)
}
