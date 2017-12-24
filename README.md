# gocrest

A hamcrest-like assertion library for Go.

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

# Matchers so far..

- EqualTo(x)
- IsNil()
- Contains(x) -- acts like containsAll
- Not(m *Matcher) -- logical not of matcher's result
- MatchesPattern(regex string) -- a string regex expression