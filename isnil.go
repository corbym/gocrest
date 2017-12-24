package gocrest

//Matches if the expected value is nil
func IsNil() *Matcher {
	return EqualTo(nil)
}