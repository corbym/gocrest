# gocrest

A hamcrest-like assertion library for Go.
## package import

{code}
import (
  gocrest "github.com/corbym/gocrest"
)
{code}
## Example:
{code}
		gocrest.AssertThat(mockTestingT, "hi", gocrest.EqualTo("hi"))
{code}
