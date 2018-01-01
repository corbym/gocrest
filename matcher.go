package gocrest

import "fmt"

type Matcher struct {
	Matches      func(actual interface{}) bool
	Describe     string
	Actual       string
	ReasonString string
}

func (matcher *Matcher) Reason(r string) *Matcher {
	matcher.ReasonString = r
	return matcher
}
func (matcher *Matcher) Reasonf(format string, args ...interface{}) *Matcher {
	return matcher.Reason(fmt.Sprintf(format, args...))
}
