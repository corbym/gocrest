package base

type Matcher struct {
	Matches  func(actual interface{}) bool
	Describe string
	Actual   string
}
