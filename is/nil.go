package is

import "gocrest"

//Matches if the expected value is nil
func Nil() *gocrest.Matcher {
	return EqualTo(nil)
}
