package then

import (
	"fmt"
	"github.com/corbym/gocrest"
	"sync"
	"time"
)

type RecordingTestingT struct {
	gocrest.TestingT
	TestOutput string
	failed     bool
}

func (t *RecordingTestingT) Logf(format string, args ...interface{}) {
	t.TestOutput = fmt.Sprintf(format, args...)
}

func (t *RecordingTestingT) Errorf(format string, args ...interface{}) {
	t.TestOutput = fmt.Sprintf(format, args...)
	t.failed = true
}
func (t *RecordingTestingT) Fail() {
	t.failed = true
}
func (t *RecordingTestingT) FailNow() {
	t.failed = true
}

type Latest struct {
	mu          sync.Mutex
	latestValue RecordingTestingT
}

func (l *Latest) Get() RecordingTestingT {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.latestValue
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
			t.Errorf(latestValue.Get().TestOutput)
			return
		case <-tick:
			tick = nil
			go func() {
				recordedTesting := RecordingTestingT{
					TestingT: t,
					failed:   false,
				}
				assertions(&recordedTesting)
				channel <- recordedTesting
			}()
		case value := <-channel:
			if !value.failed {
				return
			}
			latestValue.latestValue = value
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
