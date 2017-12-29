# gocrest

A hamcrest-like assertion library for Go. GoCrest matchers are composable, self-describing and
can be strung together in a more readable form to create flexible assertions. 

Inspired by [Hamcrest](https://github.com/hamcrest). 

## Package import

```
import (
  "github.com/corbym/gocrest/then"
  "github.com/corbym/gocrest/is"
  "github.com/corbym/gocrest/has"
)
```

## Example:
```
then.AssertThat(testing, "hi", is.EqualTo("bye"))
```

output:

```
expected: value equal to bye but was: hi
```

Composed with AllOf:

``` then.AssertThat(t, "abcdef", is.AllOf(is.ValueContaining("abc"), is.LessThan("ghi")))```

# Matchers so far..

- is.EqualTo(x)
- is.Nil()
- is.ValueContaining(expected) -- acts like containsAll
- is.Not(m *Matcher) -- logical not of matcher's result
- is.MatchForPattern(regex string) -- a string regex expression
- has.FunctionNamed(string) - checks if a Type has a function (method)
- is.AllOf(... *Matcher) - returns true if all matchers match
- is.AnyOf(... *Matcher) - return true if any matcher matches
- is.GreaterThan(expected) - checks if actual > expected
- is.LessThan(expected)
- is.Empty() - matches if the actual is "", nil or len(actual)==0