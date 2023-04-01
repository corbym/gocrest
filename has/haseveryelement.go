package has

import (
	"fmt"
	"github.com/corbym/gocrest"
)

// EveryElement Checks whether the nth element of the array/slice matches the nth expectation passed
func EveryElement[A any](expects ...*gocrest.Matcher[A]) *gocrest.Matcher[[]A] {
	match := new(gocrest.Matcher[[]A])
	match.Describe = fmt.Sprintf("elements to match %s", describe(expects, "and"))

	for _, e := range expects {
		match.AppendActual(e.Actual)
	}

	match.Matches = func(actual []A) bool {
		if len(actual) != len(expects) {
			return false
		}

		for i := 0; i < len(actual); i++ {
			result := expects[i].Matches(actual[i])
			if !result {
				return false
			}
		}

		return true
	}

	return match
}

func describe[A any](matchers []*gocrest.Matcher[A], conjunction string) string {
	var description string
	for x := 0; x < len(matchers); x++ {
		description += fmt.Sprintf("[%v]:%v", x, matchers[x].Describe)
		if x+1 < len(matchers) {
			description += fmt.Sprintf(" %s ", conjunction)
		}
	}
	return description
}
