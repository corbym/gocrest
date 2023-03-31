package is

import (
	"bytes"
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

// EqualTo checks if two values are equal. Uses DeepEqual (could be slow) or Compare for byte arrays.
// Like DeepEquals, if the types are not the same the matcher returns false.
// Returns a matcher that will return true if two values are equal.
func EqualTo[A any](expected A) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("value equal to <%v>", expected)
	match.Matches = func(actual A) bool {
		return reflect.DeepEqual(expected, actual)
	}

	return match
}
func EqualToBytes(expected []byte) *gocrest.Matcher[[]byte] {
	match := new(gocrest.Matcher[[]byte])
	match.Describe = fmt.Sprintf("value equal to <%v>", expected)
	match.Matches = func(actual []byte) bool {
		compare := bytes.Compare(expected, actual)
		return compare == 0
	}

	return match
}
