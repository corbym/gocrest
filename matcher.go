package gocrest

type Matcher struct {
	matches  func(actual interface{}) bool
	describe string
	actual   string
}
