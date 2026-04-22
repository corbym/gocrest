package is

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
	"sort"
	"strings"
)

// EqualTo checks if two values are equal. Uses DeepEqual (could be slow).
// Like DeepEquals, if the types are not the same the matcher returns false.
// Returns a matcher that will return true if two values are equal.
func EqualTo[A any](expected A) *gocrest.Matcher[A] {
	match := new(gocrest.Matcher[A])
	match.Describe = fmt.Sprintf("value equal to <%v>", expected)
	match.Matches = func(actual A) bool {
		if reflect.DeepEqual(expected, actual) {
			return true
		}
		match.Actual = diffDescription(reflect.ValueOf(expected), reflect.ValueOf(actual), "")
		return false
	}

	return match
}

// diffDescription recursively compares expected and actual reflect.Values,
// returning a human-readable description of the first differences found.
// path is the breadcrumb prefix built during recursion (e.g. ".Address").
func diffDescription(expected, actual reflect.Value, path string) string {
	// Dereference pointers only when both sides are pointers, to avoid panicking
	// on a type mismatch where one side is a pointer and the other is not.
	for expected.Kind() == reflect.Ptr && actual.Kind() == reflect.Ptr {
		if expected.IsNil() || actual.IsNil() {
			break
		}
		expected = expected.Elem()
		actual = actual.Elem()
	}

	if !expected.IsValid() && !actual.IsValid() {
		return fmt.Sprintf("%v: expected <nil>, got <nil>", path)
	}
	if !expected.IsValid() {
		return fmt.Sprintf("%v: expected <nil>, got <%v>", path, actual)
	}
	if !actual.IsValid() {
		return fmt.Sprintf("%v: expected <%v>, got <nil>", path, expected)
	}

	switch expected.Kind() {
	case reflect.Struct:
		var diffs []string
		t := expected.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}
			ev := expected.Field(i)
			av := actual.Field(i)
			if !reflect.DeepEqual(ev.Interface(), av.Interface()) {
				fieldPath := path + "." + field.Name
				diffs = append(diffs, diffDescription(ev, av, fieldPath))
			}
		}
		if len(diffs) > 0 {
			return strings.Join(diffs, ", ")
		}

	case reflect.Slice, reflect.Array:
		var diffs []string
		maxLen := expected.Len()
		if actual.Len() > maxLen {
			maxLen = actual.Len()
		}
		for i := 0; i < maxLen; i++ {
			elemPath := fmt.Sprintf("%s[%d]", path, i)
			if i >= expected.Len() {
				diffs = append(diffs, fmt.Sprintf("%s: expected <missing>, got <%v>", elemPath, actual.Index(i)))
			} else if i >= actual.Len() {
				diffs = append(diffs, fmt.Sprintf("%s: expected <%v>, got <missing>", elemPath, expected.Index(i)))
			} else {
				ev := expected.Index(i)
				av := actual.Index(i)
				if !reflect.DeepEqual(ev.Interface(), av.Interface()) {
					diffs = append(diffs, diffDescription(ev, av, elemPath))
				}
			}
		}
		if len(diffs) > 0 {
			return strings.Join(diffs, ", ")
		}

	case reflect.Map:
		// Collect diffs keyed by a sortable string label.
		type mapDiff struct {
			label string
			msg   string
		}
		var diffs []mapDiff

		// Keys present in expected — compare against actual using MapIndex directly.
		for _, k := range expected.MapKeys() {
			av := actual.MapIndex(k)
			ev := expected.MapIndex(k)
			keyPath := fmt.Sprintf("%s[%v]", path, k.Interface())
			if !av.IsValid() {
				diffs = append(diffs, mapDiff{keyPath, fmt.Sprintf("%s: expected <%v>, got <missing>", keyPath, ev)})
			} else if !reflect.DeepEqual(ev.Interface(), av.Interface()) {
				diffs = append(diffs, mapDiff{keyPath, diffDescription(ev, av, keyPath)})
			}
		}

		// Keys present only in actual (not in expected).
		for _, k := range actual.MapKeys() {
			if !expected.MapIndex(k).IsValid() {
				keyPath := fmt.Sprintf("%s[%v]", path, k.Interface())
				diffs = append(diffs, mapDiff{keyPath, fmt.Sprintf("%s: expected <missing>, got <%v>", keyPath, actual.MapIndex(k))})
			}
		}

		if len(diffs) > 0 {
			sort.Slice(diffs, func(i, j int) bool { return diffs[i].label < diffs[j].label })
			msgs := make([]string, len(diffs))
			for i, d := range diffs {
				msgs[i] = d.msg
			}
			return strings.Join(msgs, ", ")
		}
	}

	// Primitive or unexplored type: show expected vs actual directly.
	if path == "" {
		return fmt.Sprintf("%v", actual)
	}
	return fmt.Sprintf("%s: expected <%v>, got <%v>", path, expected, actual)
}
