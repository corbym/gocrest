# `then` and `by` Packages

---

## then.AssertThat(t, actual, matcher)

The core assertion function. Evaluates the matcher against the actual value and calls `t.Errorf` with a human-readable message when it does not match.

```go
then.AssertThat(t, "hello", is.EqualTo("world"))
```

```
Expected: value equal to <world>
     but: <hello>
```

Adding a reason with `.Reason(...)` prepends a message to the output:

```go
then.AssertThat(t, "hello", is.EqualTo("world").Reason("greeting must match"))
```

```
greeting must match
Expected: value equal to <world>
     but: <hello>
```

Using `.Reasonf(format, args...)` allows a formatted reason string:

```go
then.AssertThat(t, got, is.EqualTo(want).Reasonf("attempt %d failed", attempt))
```

---

## then.Eventually(t, waitFor, tick, assertions)

Repeatedly runs the assertions function on a `tick` interval until all assertions pass or the `waitFor` timeout expires.
Each run uses its own `RecordingTestingT` so failures do not immediately fail the outer test.
If the timeout is reached the last set of failures is reported.

```go
then.Eventually(t, 5*time.Second, time.Second, func(eventually gocrest.TestingT) {
    then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(42))
})
```

Output when the timeout is reached before the assertion passes:

```
Eventually Failed after 5s:
Expected: value equal to <42>
     but: <0>
```

---

## then.WithinFiveSeconds(t, assertions)

Shorthand for `then.Eventually(t, 5*time.Second, time.Second, assertions)`.

```go
then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
    then.AssertThat(eventually, by.Channelling(channel), is.EqualTo("done"))
})
```

---

## then.WithinTenSeconds(t, assertions)

Shorthand for `then.Eventually(t, 10*time.Second, time.Second, assertions)`.

```go
then.WithinTenSeconds(t, func(eventually gocrest.TestingT) {
    then.AssertThat(eventually, by.Channelling(channel), is.EqualTo("done"))
})
```

---

## by.Channelling[T](channel chan T) T

Reads one value from the channel and returns it. Intended to be used as the `actual` argument inside `then.Eventually`.

```go
ch := make(chan int, 1)
go func() { ch <- 99 }()

then.Eventually(t, 5*time.Second, time.Second, func(eventually gocrest.TestingT) {
    then.AssertThat(eventually, by.Channelling(ch), is.EqualTo(100))
})
```

```
Eventually Failed after 5s:
Expected: value equal to <100>
     but: <99>
```

---

## by.Reading(reader io.Reader, len int) []byte

Reads up to `len` bytes from the reader and returns them as a byte slice. Intended to be used as the `actual` argument inside `then.Eventually`.

```go
then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
    then.AssertThat(eventually, by.Reading(reader, 5), is.EqualTo([]byte("hello")))
})
```

```
Eventually Failed after 5s:
Expected: value equal to <[104 105]>
     but: <[0]: expected <104>, got <0>>
```

---

## by.Calling[K, T](fn func(T) K, value T) K

Calls `fn(value)` and returns the result. Use this to wrap a function call as the `actual` argument of an assertion.

```go
double := func(n int) int { return n * 2 }
then.AssertThat(t, by.Calling(double, 3), is.EqualTo(7))
```

```
Expected: value equal to <7>
     but: <6>
```
