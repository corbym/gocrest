package is

import "github.com/corbym/gocrest"

func False() *gocrest.Matcher {
	return &gocrest.Matcher{
		Describe: "is false",
		Matches: func(actual interface{}) bool {
			return actual == false
		},
	}
}
