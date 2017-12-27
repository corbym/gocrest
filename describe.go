package gocrest

import "fmt"

func describe(matchers []*Matcher, conjunction string) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += matchers[x].describe
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
