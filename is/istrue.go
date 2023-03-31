package is

import "github.com/corbym/gocrest"

// True returns true if the actual matches true
func True() *gocrest.Matcher[bool] {
	return &gocrest.Matcher[bool]{
		Describe: "is true",
		Matches: func(actual bool) bool {
			return actual == true
		},
	}
}
