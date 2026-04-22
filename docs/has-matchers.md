# `has` Package Matchers

All matchers are used with `then.AssertThat(t, actual, matcher)`.

---

## has.Key[K, V](expected K)

Checks that a map contains the given key.

```go
then.AssertThat(t, map[string]int{"a": 1}, has.Key[string, int]("b"))
```

```
Expected: map has key 'b'
     but: <map[a:1]>
```

---

## has.AllKeys[K, V](expected ...K)

Checks that a map contains **all** of the given keys.

```go
then.AssertThat(t, map[string]int{"a": 1}, has.AllKeys[string, int]("a", "b"))
```

```
Expected: map has keys '[a b]'
     but: <map[a:1]>
```

> **Note:** Map matchers often need the type parameters stated explicitly so the compiler can infer them:
> ```go
> then.AssertThat(t, map[string]bool{"hi": true}, has.AllKeys[string, bool]("hi", "bye"))
> ```

---

## has.Length[T](expected int)

Checks that a slice or array has exactly the given length.

```go
then.AssertThat(t, []int{1, 2}, has.Length[int](5))
```

```
Expected: value with length 5
     but: <length was 2>
```

---

## has.StringLength(expected int)

Checks that a string has exactly the given length in bytes.

```go
then.AssertThat(t, "hello", has.StringLength(3))
```

```
Expected: value with length 3
     but: <length was 5>
```

---

## has.MapLength[K, V](expected int)

Checks that a map has exactly the given number of entries.

```go
then.AssertThat(t, map[string]int{"a": 1, "b": 2}, has.MapLength[string, int](5))
```

```
Expected: value with length 5
     but: <length was 2>
```

---

## has.LengthMatching[A](expected *Matcher[int])

Checks that the length of a slice or array satisfies the given matcher.

```go
then.AssertThat(t, []int{1, 2, 3, 4, 5}, has.LengthMatching[int](is.LessThan(3)))
```

```
Expected: value with length value less than <3>
     but: <length was 5>
```

---

## has.StringLengthMatching(expected *Matcher[int])

Checks that the length of a string satisfies the given matcher.

```go
then.AssertThat(t, "hello world", has.StringLengthMatching(is.LessThan(3)))
```

```
Expected: value with length value less than <3>
     but: <length was 11>
```

---

## has.MapLengthMatching[K, V](expected *Matcher[int])

Checks that the number of entries in a map satisfies the given matcher.

```go
then.AssertThat(t, map[string]int{"a": 1, "b": 2}, has.MapLengthMatching[string, int](is.LessThan(1)))
```

```
Expected: value with length value less than <1>
     but: <length was 2>
```

---

## has.Prefix(expected string)

Checks that a string starts with the given prefix.

```go
then.AssertThat(t, "hello world", has.Prefix("goodbye"))
```

```
Expected: value with prefix goodbye
     but: <hello world>
```

---

## has.Suffix(expected string)

Checks that a string ends with the given suffix.

```go
then.AssertThat(t, "main.go", has.Suffix(".py"))
```

```
Expected: value with suffix .py
     but: <main.go>
```

---

## has.TypeName[A](expected string)

Checks that the actual value's type name equals the expected string.
The type name is the fully qualified name as returned by `reflect.TypeOf(actual).String()`.

```go
then.AssertThat(t, 42, has.TypeName[int]("string"))
```

```
Expected: has type <string>
     but: <int>
```

---

## has.TypeNameMatches[A](expected *Matcher[string])

Checks that the actual value's type name satisfies the given matcher.

```go
then.AssertThat(t, 42, has.TypeNameMatches[int](is.StringContaining("float")))
```

```
Expected: has type something that contains [float]
     but: <int>
```

---

## has.FunctionNamed[A](expected string)

Checks that a type has a method with the given name. Does not check parameters or return types.

```go
type MyInterface interface {
    Foo()
}
then.AssertThat(t, (*MyInterface)(nil), has.FunctionNamed[*MyInterface]("Bar"))
```

```
Expected: interface with function Bar
     but: <MyInterface{Foo()}>
```

---

## has.FieldNamed[A](expected string)

Checks that a struct has an exported field with the given name. Does not check the field type.

```go
type MyStruct struct {
    Age int
}
then.AssertThat(t, (*MyStruct)(nil), has.FieldNamed[*MyStruct]("Name"))
```

```
Expected: struct with function Name
     but: <MyStruct{Age int}>
```

---

## has.EveryElement(expects ...*Matcher[A])

Checks that the slice has the same number of elements as matchers provided, and that each element at index `i` satisfies the matcher at index `i`.

```go
then.AssertThat(t, []int{1, 5, 3}, has.EveryElement(
    is.EqualTo(1),
    is.GreaterThan(10),
    is.EqualTo(3),
))
```

```
Expected: elements to match [0]:value equal to <1> and [1]:value greater than <10> and [2]:value equal to <3>
     but: <>
```

---

## has.StructWithValues[A, B](expects StructMatchers[B])

Checks that the named fields of a struct each satisfy their corresponding matcher.
Only the fields named in the `StructMatchers` map are checked — other fields are ignored.
Automatically dereferences pointers. Panics if a named field does not exist or is unexported.

```go
type Person struct {
    Name string
    Age  int
}
then.AssertThat(t, Person{Name: "Alice", Age: 25},
    has.StructWithValues[Person, string](has.StructMatchers[string]{
        "Name": is.EqualTo("Bob"),
    }),
)
```

```
Expected: struct values to match {"Name": value equal to <Bob>}
     but: <>
```
