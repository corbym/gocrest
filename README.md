# gocrest

A hamcrest-like assertion library for Go. GoCrest matchers are composable, self-describing and
can be strung together in a more readable form to create flexible assertions. 

Inspired by [Hamcrest](https://github.com/hamcrest). 

[![Build](https://github.com/corbym/gocrest/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/corbym/gocrest/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/gocrest)](https://goreportcard.com/report/github.com/corbym/gocrest)
[![GoDoc](https://godoc.org/github.com/corbym/gocrest?status.svg)](http://godoc.org/github.com/corbym/gocrest)
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
then.AssertThat(t, "abcdef", is.AllOf(is.ValueContaining("abc"), is.LessThan("ghi")))
```

Asynchronous Matching (v1.0.8 onwards):

```go
//Reader
then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
	then.AssertThat(eventually, by.Reading(slowReader, 1024), is.EqualTo([]byte("abcdefghijklmnopqrstuv")))
})
```
```go
//channels
then.Eventually(t, time.Second*5, time.Second, func(eventually gocrest.TestingT) {
	then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
})
```
```go
// multiple assertions
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
* `has.Length()` is likewise pernickety about types being explicit, mainly because it works on both strings and arrays. It needs to know both the type of the array and the array/string type. Confused? me too.
* `is.LessThan()` and `is.GreaterThan()` (and by extension `is.GreaterThanOrEqualTo` and `is.LessThanOrEqualTo`) no longer work on complex types. This is because the complex types do not support the comparison operators (yet, somehow, they could be compared by reflection 🤷 )

See the matcher_test.go file for full usage.

# Matchers so far..

- is.EqualTo(x)
- is.EqualToIgnoringWhitespace(string) - compares two strings without comparing their whitespace characters.
- is.Nil() - value must be nil
- is.StringContaining(expected) -- acts like containsAll
- is.MapContaining(expected) -- acts like containsAll
- is.MapContainingValues(expected) -- acts like containsAll
- is.MapMatchingValues(expected) -- acts like containsAll
- is.ArrayContaining(expected) -- acts like containsAll
- is.ArrayMatching(expected) -- acts like containsAll
- is.Not(m *Matcher) -- logical not of matcher's result
- is.MatchForPattern(regex string) -- a string regex expression
- has.FunctionNamed(string x) - checks if an interface has a function (method)
- has.FieldNamed(string x) - checks if a struct has a field named x
- is.AllOf(... *Matcher) - returns true if all matchers match
- is.AnyOf(... *Matcher) - return true if any matcher matches
- is.GreaterThan(expected) - checks if actual > expected
- is.LessThan(expected)
- is.Empty() - matches if the actual is "", nil or len(actual)==0
- is.LessThan(x)
- is.LessThanOrEqualTo(x)
- is.GreaterThan(x)
- is.GreaterThanOrEqualTo(x)
- has.Length(x) - matcher if given value (int or matcher) matches the len of the given
- has.Prefix(x) - string starts with x
- has.Suffix(x) - string ends with x
- has.Key(x) - map has key x
- has.AllKeys(T x, T y) (or has.AllKeys([]T{x,y})) - finds key of type T in map
- has.EveryElement(x1...xn) - checks if actual[i] matches corresponding expectation (x[i])
- has.StructWithValues(map[string]*gocrest.Matcher) - checks if actual[key] matches corresponding expectation (x[key])

For more comprehensive documentation see [godoc](http://godoc.org/github.com/corbym/gocrest).
