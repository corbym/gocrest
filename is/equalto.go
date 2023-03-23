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
func EqualTo(expected interface{}) *gocrest.Matcher {
	match := new(gocrest.Matcher)
	match.Describe = fmt.Sprintf("value equal to <%v>", expected)
	match.Matches = func(actual interface{}) bool {
		switch actual.(type) {
		case []byte:
			expectedBytes := expected.([]byte)
			actualBytes := actual.([]byte)
			compare := bytes.Compare(expectedBytes, actualBytes)
			return compare == 0
		default:
			return reflect.DeepEqual(expected, actual)
		}
	}

	return match
}
