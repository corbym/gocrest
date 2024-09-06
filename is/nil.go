package is

import (
	"github.com/corbym/gocrest"
)

var description = "value that is <nil>"

// Nil matches if the actual error value is nil
func Nil() *gocrest.Matcher[error] {
	match := new(gocrest.Matcher[error])
	match.Describe = description
	match.Matches = func(actual error) bool {
		return actual == nil
	}
	return match
}

// NilArray matches if the actual array value is nil
func NilArray[A any]() *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = description
	match.Matches = func(actual []A) bool {
		return actual == nil
	}
	return match
}

// NilMap matches if the actual map value is nil
func NilMap[K comparable, V any]() *gocrest.Matcher[map[K]V] {
	match := new(gocrest.Matcher[map[K]V])
	match.Describe = description
	match.Matches = func(actual map[K]V) bool {
		return actual == nil
	}
	return match
}

// NilPtr matches if the actual pointer to T is nil
func NilPtr[T any]() *gocrest.Matcher[*T] {
	match := new(gocrest.Matcher[*T])
	match.Describe = description
	match.Matches = func(actual *T) bool {
		return actual == nil
	}
	return match
}
