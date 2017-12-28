# gocrest

A hamcrest-like assertion library for Go. GoCrest matchers are composable and
can be strung together in a more readable form to create flexible assertions. 

Inspired by [Hamcrest](https://github.com/hamcrest). 

## Package import

```
import (
  gocrest "github.com/corbym/gocrest"
)
```

## Example:
```
gocrest.AssertThat(testing, "hi", gocrest.EqualTo("bye"))
```

output:

```
expected: value equal to bye but was: hi
```

Composed with AllOf:

``` gocrest.AssertThat(t, "abcdef", gocrest.AllOf(gocrest.Contains("abc"), gocrest.LessThan("ghi")))```

# Matchers so far..

- EqualTo(x)
- IsNil()
- Contains(expected) -- acts like containsAll
- Not(m *Matcher) -- logical not of matcher's result
- MatchesPattern(regex string) -- a string regex expression
- HasFunction(string) - checks if a Type has a function (method)
- AllOf(... *Matcher) - returns true if all matchers match
- AnyOf(... *Matcher) - return true if any matcher matches
- GreaterThan(expected) - checks if actual > expected
- LessThan(expected)