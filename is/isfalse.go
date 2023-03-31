package is

import "github.com/corbym/gocrest"

// False returns true if the actual matches false. Confusing but true.
func False() *gocrest.Matcher[bool] {
	return &gocrest.Matcher[bool]{
		Describe: "is false",
		Matches: func(actual bool) bool {
			return actual == false
		},
	}
}
