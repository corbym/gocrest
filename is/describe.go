package is

import (
	"fmt"
	"gocrest/base"
)

func describe(matchers []*base.Matcher, conjunction string) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += matchers[x].Describe
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
