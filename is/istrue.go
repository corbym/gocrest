package is

import "github.com/corbym/gocrest"

func True() *gocrest.Matcher {
	return &gocrest.Matcher{
		Describe: "is true",
		Matches: func(actual interface{}) bool {
			return actual == true
		},
	}
}
