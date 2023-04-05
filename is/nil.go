package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

var description = "value that is <nil>"

// Nil matches if the actual value is nil
func Nil() *gocrest.Matcher[error] {
	match := new(gocrest.Matcher[error])
	match.Describe = description
	match.Matches = func(actual error) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		return actual == nil
	}
	return match
}
func NilArray[A any]() *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = description
	match.Matches = func(actual []A) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		return actual == nil
	}
	return match
}
func NilMap[K comparable, V any]() *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = description
	match.Matches = func(actual map[K]V) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		return actual == nil
	}
	return match
}
func NilPtr[T any]() *gocrest.Matcher[*T] {
	match := new(gocrest.Matcher[*T])
	match.Describe = description
	match.Matches = func(actual *T) bool {
		match.Actual = fmt.Sprintf("%v", actual)
		return actual == nil
	}
	return match
}
