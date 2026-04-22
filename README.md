# gocrest

A hamcrest-like assertion library for Go. GoCrest matchers are composable, self-describing and
can be strung together in a more readable form to create flexible assertions. 

Inspired by [Hamcrest](https://github.com/hamcrest). 

[![Build](https://github.com/corbym/gocrest/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/corbym/gocrest/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/gocrest)](https://goreportcard.com/report/github.com/corbym/gocrest)
[![GoDoc](https://pkg.go.dev/badge/github.com/corbym/gocrest.svg)](https://pkg.go.dev/github.com/corbym/gocrest)
[![Coverage Status](https://coveralls.io/repos/github/corbym/gocrest/badge.svg?branch=master)](https://coveralls.io/github/corbym/gocrest?branch=master)
## Package import

```
import (
  "github.com/corbym/gocrest/by"
  "github.com/corbym/gocrest/then"
  "github.com/corbym/gocrest/is"
  "github.com/corbym/gocrest/has"
)
```

## Example:
```go
then.AssertThat(testing, "hi", is.EqualTo("bye").Reason("we are going"))
```

output:

```
we are going
Expected: value equal to <bye>
     but: <hi>
```

Composed with AllOf:

```go
then.AssertThat(t, "abcdef", is.AllOf(is.StringContaining("abc"), is.LessThan("ghi")))
```

Asynchronous Matching (v1.0.8 onwards):

```go
// Reader
then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
	then.AssertThat(eventually, by.Reading(slowReader, 1024), is.EqualTo([]byte("abcdefghijklmnopqrstuv")))
})
```
```go
// Channels
then.Eventually(t, time.Second*5, time.Second, func(eventually gocrest.TestingT) {
	then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
})
```
```go
// Calling a function
then.AssertThat(t, by.Calling(myFunc, inputValue), is.EqualTo(expectedOutput))
```
```go
// Multiple assertions
then.WithinTenSeconds(t, func(eventually gocrest.TestingT) {
	then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
	then.AssertThat(eventually, by.Channelling(channelTwo), is.EqualTo("11").Reason("This is unreachable"))
})
```
# v.1.1.0 - generics

Changes all the matchers to use generics instead of reflection. Some still use a bit of reflection, e.g. TypeName etc.

## Other major changes:

* ValueContaining has been split into StringContaining, MapContaining, MapContainingValues, MapMatchingValues, ArrayContaining and ArrayMatching.
* No longer panics with unknown types, as types will fail at compile time.
Some idiosyncrasies with the generic types do exist, but this is language specific;

* Map matchers generally need to know the type of the map key values explicitly or the compiler will complain, e.g.
```
then.AssertThat(testing, map[string]bool{"hi": true, "bye": true}, has.AllKeys[string, bool]("hi", "bye"))
```
* Length matchers are type-specific: use `has.Length[T]()` for arrays/slices, `has.StringLength()` for strings, and `has.MapLength[K, V]()` for maps. Matcher variants (`has.LengthMatching`, `has.StringLengthMatching`, `has.MapLengthMatching`) accept a `*Matcher[int]` instead of a plain int.
* `is.LessThan()` and `is.GreaterThan()` (and by extension `is.GreaterThanOrEqualTo` and `is.LessThanOrEqualTo`) no longer work on complex types. This is because the complex types do not support the comparison operators (yet, somehow, they could be compared by reflection 🤷 )

See the matcher_test.go file for full usage.

# Matchers so far..

- is.EqualTo(x)
- is.EqualToIgnoringWhitespace(string) - compares two strings without comparing their whitespace characters.
- is.Nil() - error value must be nil
- is.NilArray() - array/slice value must be nil
- is.NilMap() - map value must be nil
- is.NilPtr() - pointer value must be nil
- is.True() - boolean value must be true
- is.False() - boolean value must be false
- is.StringContaining(expected) -- acts like containsAll
- is.MapContaining(expected) -- acts like containsAll
- is.MapContainingValues(expected) -- acts like containsAll
- is.MapMatchingValues(expected) -- acts like containsAll
- is.ArrayContaining(expected) -- acts like containsAll
- is.ArrayMatching(expected) -- acts like containsAll
- is.Not(m *Matcher) -- logical not of matcher's result
- is.MatchForPattern(regex string) -- a string regex expression
- is.AllOf(... *Matcher) - returns true if all matchers match
- is.AnyOf(... *Matcher) - returns true if any matcher matches
- is.GreaterThan(expected) - checks if actual > expected
- is.GreaterThanOrEqualTo(expected)
- is.LessThan(expected) - checks if actual < expected
- is.LessThanOrEqualTo(expected)
- is.Empty() - matches if the actual slice has len == 0
- is.EmptyString() - matches if the actual string is ""
- is.EmptyMap() - matches if the actual map has len == 0
- has.FunctionNamed(x string) - checks if an interface has a function (method) named x
- has.FieldNamed(x string) - checks if a struct has a field named x
- has.Length[T](expected int) - matches if the length of an array/slice equals expected
- has.StringLength(expected int) - matches if the length of a string equals expected
- has.MapLength(expected int) - matches if the length of a map equals expected
- has.LengthMatching[A](expected *Matcher[int]) - matches if the length of an array/slice satisfies expected matcher
- has.StringLengthMatching(expected *Matcher[int]) - matches if the length of a string satisfies expected matcher
- has.MapLengthMatching(expected *Matcher[int]) - matches if the length of a map satisfies expected matcher
- has.Prefix(x) - string starts with x
- has.Suffix(x) - string ends with x
- has.Key(x) - map has key x
- has.AllKeys(x ...K) - matches if all given keys are present in the map
- has.TypeName(expected string) - matches if the actual value's type name equals expected
- has.TypeNameMatches(expected *Matcher[string]) - matches if the actual value's type name satisfies expected matcher
- has.EveryElement(x1...xn) - checks if actual[i] matches corresponding expectation (x[i])
- has.StructWithValues(expects StructMatchers[B]) - checks if actual struct fields match their corresponding matchers

## Matcher documentation

- [is package matchers](docs/is-matchers.md) — equality, nil, boolean, string/array/map containment, comparisons, patterns, AllOf/AnyOf, Not, Empty
- [has package matchers](docs/has-matchers.md) — keys, length, prefix/suffix, type name, struct fields and methods, element and struct value matching
- [then and by packages](docs/then-and-by.md) — AssertThat, Eventually, WithinFiveSeconds, WithinTenSeconds, Channelling, Reading, Calling

For generated API documentation see [pkg.go.dev](https://pkg.go.dev/github.com/corbym/gocrest).
