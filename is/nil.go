package is

import "gocrest/base"

//Matches if the expected value is nil
func Nil() *base.Matcher {
	return EqualTo(nil)
}
