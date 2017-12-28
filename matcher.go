package gocrest

type Matcher struct {
	Matches  func(actual interface{}) bool
	Describe string
	Actual   string
}
