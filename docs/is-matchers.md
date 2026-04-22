# `is` Package Matchers

All matchers are used with `then.AssertThat(t, actual, matcher)`.

---

## is.EqualTo(expected)

Checks equality using `reflect.DeepEqual`. Works with any comparable type.
For complex types (structs, slices, maps) the mismatch output shows a field-level diff.

```go
then.AssertThat(t, "hello", is.EqualTo("world"))
```

```
Expected: value equal to <world>
     but: <hello>
```

Struct diff example:

```go
type Person struct { Name string; Age int }
then.AssertThat(t, Person{Name: "Alice", Age: 30}, is.EqualTo(Person{Name: "Bob", Age: 30}))
```

```
Expected: value equal to <{Bob 30}>
     but: <.Name: expected <Bob>, got <Alice>>
```

Slice diff example:

```go
then.AssertThat(t, []string{"actual"}, is.EqualTo([]string{"expected"}))
```

```
Expected: value equal to <[expected]>
     but: <[0]: expected <expected>, got <actual>>
```

---

## is.EqualToIgnoringWhitespace(expected)

Compares two strings after stripping all whitespace. Tabs, spaces and newlines are all ignored.

```go
then.AssertThat(t, "a   b\tc", is.EqualToIgnoringWhitespace("xyz"))
```

```
Expected: ignoring whitespace value equal to <xyz>
     but: <a   b	c>
```

---

## is.Nil()

Checks that an `error` value is nil.

```go
then.AssertThat(t, errors.New("something went wrong"), is.Nil())
```

```
Expected: value that is <nil>
     but: <something went wrong>
```

---

## is.NilArray[A]()

Checks that a slice/array value is nil (not just empty).

```go
then.AssertThat(t, []string{"a", "b"}, is.NilArray[string]())
```

```
Expected: value that is <nil>
     but: <[a b]>
```

---

## is.NilMap[K, V]()

Checks that a map value is nil (not just empty).

```go
then.AssertThat(t, map[string]int{"a": 1}, is.NilMap[string, int]())
```

```
Expected: value that is <nil>
     but: <map[a:1]>
```

---

## is.NilPtr[T]()

Checks that a pointer value is nil.

```go
s := "hello"
then.AssertThat(t, &s, is.NilPtr[string]())
```

```
Expected: value that is <nil>
     but: <0xc0000b4020>
```

---

## is.True()

Checks that a boolean value is `true`.

```go
then.AssertThat(t, false, is.True())
```

```
Expected: is true
     but: <false>
```

---

## is.False()

Checks that a boolean value is `false`.

```go
then.AssertThat(t, true, is.False())
```

```
Expected: is false
     but: <true>
```

---

## is.Not(matcher)

Negates the result of the given matcher.

```go
then.AssertThat(t, "hello", is.Not(is.EqualTo("hello")))
```

```
Expected: not(value equal to <hello>)
     but: <hello>
```

---

## is.StringContaining(expected ...string)

Checks that a string contains all the given substrings (acts like "contains all").

```go
then.AssertThat(t, "hello world", is.StringContaining("foo", "bar"))
```

```
Expected: something that contains [foo bar]
     but: <hello world>
```

---

## is.MapContaining(expected map[K]V)

Checks that a map contains all the given key-value pairs.

```go
then.AssertThat(t, map[string]int{"a": 1}, is.MapContaining(map[string]int{"b": 2}))
```

```
Expected: something that contains map[b:2]
     but: <map[a:1]>
```

---

## is.MapContainingValues(expected ...V)

Checks that a map contains all the given values, regardless of which key they are stored under.

```go
then.AssertThat(t, map[string]int{"a": 1}, is.MapContainingValues[string](2, 3))
```

```
Expected: something that contains [2 3]
     but: <map[a:1]>
```

---

## is.MapMatchingValues(expected ...*Matcher[V])

Checks that a map contains values satisfying every given matcher.

```go
then.AssertThat(t, map[string]int{"a": 1}, is.MapMatchingValues[string](is.GreaterThan(10)))
```

```
Expected: value greater than <10>
     but: <map[a:1]>
```

---

## is.ArrayContaining(expected ...A)

Checks that a slice contains all given elements (acts like "contains all").

```go
then.AssertThat(t, []int{1, 2, 3}, is.ArrayContaining(4, 5))
```

```
Expected: something that contains <4> and <5>
     but: <[1 2 3]>
```

---

## is.ArrayMatching(expected ...*Matcher[A])

Checks that a slice contains elements satisfying every given matcher.

```go
then.AssertThat(t, []int{1, 2, 3}, is.ArrayMatching(is.GreaterThan(10)))
```

```
Expected: something that contains <value greater than <10>>
     but: <[1 2 3]>
```

---

## is.MatchForPattern(regex)

Checks that a string matches the given regular expression.

```go
then.AssertThat(t, "hello", is.MatchForPattern("[0-9]+"))
```

```
Expected: a value that matches pattern [0-9]+
     but: <hello>
```

---

## is.AllOf(matchers ...)

Passes only if **all** given matchers match. On failure, lists only the failing matchers.

```go
then.AssertThat(t, "hello", is.AllOf(is.StringContaining("foo"), is.StringContaining("bar")))
```

```
Expected: something that contains [foo] and something that contains [bar]
     but: actual <hello>
```

---

## is.AnyOf(matchers ...)

Passes if **at least one** of the given matchers matches. Fails when none do.

```go
then.AssertThat(t, "hello", is.AnyOf(is.EqualTo("foo"), is.EqualTo("bar")))
```

```
Expected: any of (value equal to <foo> or value equal to <bar>)
     but: actual <hello>
```

---

## is.GreaterThan(expected)

Checks that the actual value is greater than expected. Works on all numeric types and strings (lexicographic comparison).

```go
then.AssertThat(t, 3, is.GreaterThan(5))
```

```
Expected: value greater than <5>
     but: <3>
```

---

## is.GreaterThanOrEqualTo(expected)

Checks that the actual value is greater than or equal to expected.

```go
then.AssertThat(t, 3, is.GreaterThanOrEqualTo(5))
```

```
Expected: any of (value greater than <5> or value equal to <5>)
     but: actual <3>
```

---

## is.LessThan(expected)

Checks that the actual value is less than expected.

```go
then.AssertThat(t, 10, is.LessThan(5))
```

```
Expected: value less than <5>
     but: <10>
```

---

## is.LessThanOrEqualTo(expected)

Checks that the actual value is less than or equal to expected.

```go
then.AssertThat(t, 10, is.LessThanOrEqualTo(5))
```

```
Expected: any of (value less than <5> or value equal to <5>)
     but: actual <10>
```

---

## is.Empty[A]()

Checks that a slice has length 0.

```go
then.AssertThat(t, []string{"hello"}, is.Empty[string]())
```

```
Expected: empty value
     but: <[hello]>
```

---

## is.EmptyString()

Checks that a string is `""`.

```go
then.AssertThat(t, "hello", is.EmptyString())
```

```
Expected: empty value
     but: <hello>
```

---

## is.EmptyMap[K, V]()

Checks that a map has length 0.

```go
then.AssertThat(t, map[string]int{"a": 1}, is.EmptyMap[string, int]())
```

```
Expected: empty value
     but: <map[a:1]>
```
