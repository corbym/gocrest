package gocrest

//Returns a matcher that returns logical not of the matcher given
func Not(matcher *Matcher) *Matcher {
	match := new(Matcher)
	match.describe = "not(" + matcher.describe + ")"
	match.matches = func(actual interface{}) bool {
		return !matcher.matches(actual)
	}
	return match
}
