package is

import "github.com/corbym/gocrest"

// True returns true if the actual matches true
func True[A bool]() *gocrest.Matcher[A] {
	return &gocrest.Matcher[A]{
		Describe: "is true",
		Matches: func(actual A) bool {
			return actual == true
		},
	}
}
