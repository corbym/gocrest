package has

import (
	"github.com/corbym/gocrest"
	"reflect"
)

// TypeName returns true if the expected matches the actual Type's Name.
// E.g. has.TypeName("pkg.Type") would be true with instance of `type Type struct{}` in package name 'pkg'.
func TypeName[A any](expected string) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Matches = func(actual A) bool {
		actualTypeName := reflect.TypeOf(actual).String()
		matcher.Actual = actualTypeName
		matcher.Describe = "has type "
		matcher.Describe += "<" + expected + ">"
		return actualTypeName == expected

	}
	return matcher
}

// TypeNameMatches returns true if the expected matches the actual Type's Name using the given matcher.
// E.g. has.TypeName(is.EqualTo("pkg.Type")) would be true with instance of `type Type struct{}` in package name 'pkg'.
func TypeNameMatches[A any](expected *gocrest.Matcher[string]) *gocrest.Matcher[A] {
	matcher := new(gocrest.Matcher[A])
	matcher.Matches = func(actual A) bool {
		actualTypeName := reflect.TypeOf(actual).String()
		matcher.Actual = actualTypeName
		matcher.Describe = "has type "

		matches := expected.Matches(actualTypeName)
		matcher.AppendActual(expected.Actual)
		matcher.Describe += expected.Describe
		return matches
	}
	return matcher
}
