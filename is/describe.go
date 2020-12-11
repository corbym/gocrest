package is

import (
	"fmt"
	"github.com/corbym/gocrest"
)

func describe(conjunction string, matchers []*gocrest.Matcher) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += matchers[x].Describe
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
