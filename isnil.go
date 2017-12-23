package gocrest


func IsNil() *Matcher {
	return EqualTo(nil)
}