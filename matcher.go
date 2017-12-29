package gocrest

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
