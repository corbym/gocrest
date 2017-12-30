package gocrest_test

import (
	"testing"
	"strings"
	"gocrest/then"
	"gocrest/is"
	"gocrest/has"
	"gocrest"
)

var stubTestingT *StubTestingT

func init() {
	stubTestingT = new(StubTestingT)
}

func TestHasLengthMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: "", expected: 0, shouldFail: false},
		{actual: "a", expected: 1, shouldFail: false},
		{actual: "1", expected: 1, shouldFail: false},
		{actual: []string{}, expected: 0, shouldFail: false},
		{actual: []string{"foo"}, expected: 1, shouldFail: false},
		{actual: []string{"foo"}, expected: 2, shouldFail: true},
		{actual: []string{"foo", "bar"}, expected: 2, shouldFail: false},
		{actual: map[string]bool{"hello": true}, expected: 1, shouldFail: false},
		{actual: map[string]bool{"helloa": true}, expected: is.LessThan(1), shouldFail: true},
		{actual: map[string]bool{"hellob": true}, expected: is.LessThanOrEqualTo(2), shouldFail: false},
		{actual: map[string]bool{"helloc": true}, expected: is.GreaterThan(2), shouldFail: true},
		{actual: map[string]bool{"hellod": true}, expected: is.GreaterThanOrEqualTo(1), shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.Length(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoValuesAreEqualOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: true},
		{actual: "hi", expected: "bees", shouldFail: true},
		{actual: map[string]bool{"hello": true}, expected: map[string]bool{"hello": true}, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.EqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, EqualTo(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestEmptyStringIsEmptyPasses(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		shouldFail bool
	}{
		{actual: "hi", shouldFail: true},
		{actual: nil, shouldFail: false},
		{actual: "", shouldFail: false},
		{actual: map[string]bool{"hello": true}, shouldFail: true},
		{actual: map[string]bool{}, shouldFail: false},
		{actual: []string{}, shouldFail: false},
		{actual: []string{"boo"}, shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.Empty())
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, Empty()) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoValuesAreGreaterThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 2, expected: 1, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: float32(1.12), expected: float32(1.0), shouldFail: false},
		{actual: float64(1.24), expected: float64(1.0), shouldFail: false},
		{actual: uint(3), expected: uint(1), shouldFail: false},
		{actual: uint16(4), expected: uint16(1), shouldFail: false},
		{actual: uint32(6), expected: uint32(1), shouldFail: false},
		{actual: uint64(7), expected: uint64(1), shouldFail: false},
		{actual: uint64(8), expected: uint64(1), shouldFail: false},
		{actual: int16(9), expected: int16(1), shouldFail: false},
		{actual: int32(10), expected: int32(1), shouldFail: false},
		{actual: int64(11), expected: int64(1), shouldFail: false},
		{actual: int64(12), expected: int64(1), shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: true},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.GreaterThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, GreaterThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatHasLengthFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, "a", has.Length(2))
	if !strings.Contains(stubTestingT.MockTestOutput, "value with length 2") {
		testing.Errorf("did not get expected description, got: %s", stubTestingT.MockTestOutput)
	}
}

func TestAssertThatTwoValuesAreGreaterThanOrEqualToOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 2, expected: 1, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: false},
		{actual: 1.0, expected: 1.12, shouldFail: true},
		{actual: float32(1.12), expected: float32(1.0), shouldFail: false},
		{actual: float64(1.24), expected: float64(1.0), shouldFail: false},
		{actual: float64(0.5), expected: float64(1.0), shouldFail: true},
		{actual: uint(1), expected: uint(1), shouldFail: false},
		{actual: uint(3), expected: uint(1), shouldFail: false},
		{actual: uint16(4), expected: uint16(1), shouldFail: false},
		{actual: uint32(6), expected: uint32(1), shouldFail: false},
		{actual: uint64(7), expected: uint64(1), shouldFail: false},
		{actual: uint64(8), expected: uint64(1), shouldFail: false},
		{actual: int16(9), expected: int16(1), shouldFail: false},
		{actual: int32(10), expected: int32(1), shouldFail: false},
		{actual: int64(11), expected: int64(1), shouldFail: false},
		{actual: int64(12), expected: int64(1), shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: true},
		{actual: "aaa", expected: "aaa", shouldFail: false},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.GreaterThanOrEqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, GreaterThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoValuesAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: float32(1.0), expected: float32(1.12), shouldFail: false},
		{actual: float64(1.0), expected: float64(1.24), shouldFail: false},
		{actual: uint(0), expected: uint(3), shouldFail: false},
		{actual: uint16(1), expected: uint16(4), shouldFail: false},
		{actual: uint32(1), expected: uint32(6), shouldFail: false},
		{actual: uint64(1), expected: uint64(7), shouldFail: false},
		{actual: uint64(1), expected: uint64(8), shouldFail: false},
		{actual: int16(1), expected: int16(9), shouldFail: false},
		{actual: int32(1), expected: int32(10), shouldFail: false},
		{actual: int64(1), expected: int64(11), shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: true},
		{actual: "aaa", expected: "aaa", shouldFail: true},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: true}, // cannot compare complex types, so fails
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoValuesAreLessThanOrEqualToPassesOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: false},
		{actual: 2.3, expected: 1.12, shouldFail: true},
		{actual: float32(1.0), expected: float32(1.12), shouldFail: false},
		{actual: float64(1.0), expected: float64(1.24), shouldFail: false},
		{actual: float64(1.0), expected: float64(1.0), shouldFail: false},
		{actual: float64(1.0), expected: float64(0.5), shouldFail: true},
		{actual: uint(0), expected: uint(0), shouldFail: false},
		{actual: uint16(1), expected: uint16(4), shouldFail: false},
		{actual: uint32(1), expected: uint32(6), shouldFail: false},
		{actual: uint64(1), expected: uint64(7), shouldFail: false},
		{actual: uint64(1), expected: uint64(8), shouldFail: false},
		{actual: int16(1), expected: int16(9), shouldFail: false},
		{actual: int32(1), expected: int32(10), shouldFail: false},
		{actual: int64(1), expected: int64(11), shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: true},
		{actual: "aaa", expected: "aaa", shouldFail: false},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThanOrEqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestNotReturnsTheOppositeOfGivenMatcher(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.Not(is.EqualTo(1)))
	if !stubTestingT.HasFailed() {
		testing.Error("Not(EqualTo) did not fail the test")
	}
}

func TestIsNilMatches(testing *testing.T) {
	then.AssertThat(testing, nil, is.Nil())
}

func TestIsNilFails(testing *testing.T) {
	then.AssertThat(stubTestingT, 2, is.Nil())
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoStringArraysTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	then.AssertThat(stubTestingT, actualList, is.ValueContaining(expectedList))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoIntArraysTest(testing *testing.T) {
	actualList := []int{12, 13}
	expectedList := []int{14, 15}
	then.AssertThat(stubTestingT, actualList, is.ValueContaining(expectedList))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Foo"
	then.AssertThat(testing, actualList, is.ValueContaining(expected))
}

func TestContainsFailsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Moo"
	then.AssertThat(stubTestingT, actualList, is.ValueContaining(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForSlice(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Baz", "Bing", "Bong"}
	then.AssertThat(testing, actualList[2:2], is.ValueContaining(expected[2:2]))
}

func TestContainsForList(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Boom", "Bong", "Bar"}
	then.AssertThat(testing, actualList, is.ValueContaining(expected))
}

func TestMapContainsMap(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	expected := map[string]string{
		"bing": "boop",
	}

	then.AssertThat(testing, actualList, is.ValueContaining(expected))
}

func TestStringContainsString(testing *testing.T) {
	actualList := "abcd"
	expected := "bc"
	then.AssertThat(testing, actualList, is.ValueContaining(expected))
}

func TestMatchesPatternMatchesString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^blarney.*"
	then.AssertThat(testing, actual, is.MatchForPattern(expected))
}

func TestMatchesPatternDoesNotMatchString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^123.?.*"
	then.AssertThat(stubTestingT, actual, is.MatchForPattern(expected))
	if !stubTestingT.HasFailed() {
		testing.Error("did not fail test")
	}
}

func TestHasPrefixPasses(testing *testing.T) {
	actual := "blarney stone"
	expected := "blarney"
	then.AssertThat(testing, actual, has.Prefix(expected))
}

func TestHasPrefixDoesNotStartWithString(testing *testing.T) {
	actual := "blarney stone"
	expected := "123"

	then.AssertThat(stubTestingT, actual, has.Prefix(expected))
	if !stubTestingT.HasFailed() {
		testing.Error("did not fail test")
	}
}

func TestHasSuffixPasses(testing *testing.T) {
	actual := "blarney stone"
	expected := "stone"
	then.AssertThat(testing, actual, has.Suffix(expected))
}

func TestHasSuffixDoesNotEndWithString(testing *testing.T) {
	actual := "blarney stone"
	expected := "123"

	then.AssertThat(stubTestingT, actual, has.Suffix(expected))
	if !stubTestingT.HasFailed() {
		testing.Error("did not fail test")
	}
}

func TestHasFunctionPasses(testing *testing.T) {
	type MyType interface {
		N() int
		f() string
	}
	actual := new(MyType)
	expected := "f"
	then.AssertThat(testing, actual, has.FunctionNamed(expected))
}

func TestHasFunctionDoesNotPass(testing *testing.T) {
	type MyType interface {
		F() string
	}
	actual := new(MyType)
	expected := "E"
	then.AssertThat(stubTestingT, actual, has.FunctionNamed(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestHasFieldNamedPasses(testing *testing.T) {
	type T struct {
		f int
	}
	actual := new(T)
	expected := "f"
	then.AssertThat(testing, actual, has.FieldNamed(expected))
}

func TestHasFieldDoesNotPass(testing *testing.T) {
	type T struct {
		F int
	}
	actual := new(T)
	expected := "E"
	then.AssertThat(stubTestingT, actual, has.FieldNamed(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAllOfMatches(testing *testing.T) {
	actual := "abcdef"
	then.AssertThat(testing, actual, is.AllOf(is.EqualTo("abcdef"), is.ValueContaining("e")))
}

func TestAllOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AllOf(is.EqualTo("abc"), is.ValueContaining("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAnyOfMatches(testing *testing.T) {
	actual := "abcdef"
	then.AssertThat(testing, actual, is.AnyOf(is.EqualTo("abcdef"), is.ValueContaining("g")))
}

func TestAnyOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AnyOf(is.EqualTo("efg"), is.ValueContaining("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestHasKeyMatches(testing *testing.T) {
	type T struct{}
	expectedT := new(T)
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: map[string]bool{"hi": true}, expected: "hi", shouldFail: false},
		{actual: map[*T]bool{expectedT: true}, expected: "hi", shouldFail: true},
		{actual: map[*T]bool{expectedT: true}, expected: expectedT, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.Key(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("unexpected result HasKey: wanted fail was %v but failed %v", test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestHasKeysMatches(testing *testing.T) {
	type T struct{}
	expectedT := new(T)
	secondExpectedT := new(T)
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: map[string]bool{"hi": true, "bye": true}, expected: []string{"hi", "bye"}, shouldFail: false},
		{actual: map[*T]bool{expectedT: true, secondExpectedT: true}, expected: []*T{expectedT, secondExpectedT}, shouldFail: false},
		{actual: map[*T]bool{expectedT: true}, expected: "foo", shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.AllKeys(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("unexpected result HasKeys(%v): wanted fail was %v but failed %v", test.actual, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestHasKeysWithVariadic(testing *testing.T) {
	actual := map[string]bool{"hi": true, "bye": false}
	then.AssertThat(testing, actual, has.AllKeys("hi", "bye"))
}

func TestMatcherDescription(testing *testing.T) {
	var equalsItems = []struct {
		description string
		actual      interface{}
		matcher     *gocrest.Matcher
		expected    string
	}{
		{description: "EqualTo.Reasonf", actual: 1, matcher: is.EqualTo(2).Reasonf("arithmetic %s is wrong", "foo"), expected: "arithmetic foo is wrong"},
		{description: "EqualTo.Reason", actual: 1, matcher: is.EqualTo(2).Reason("arithmetic is wrong"), expected: "arithmetic is wrong\nExpected: value equal to 2\n     but: 1\n"},
		{description: "Empty", actual: map[string]bool{"foo": true}, matcher: is.Empty(), expected: "empty value"},
		{description: "GreaterThan", actual: 1, matcher: is.GreaterThan(2), expected: "value greater than 2"},
		{description: "GreaterThanOrEqual", actual: 1, matcher: is.GreaterThanOrEqualTo(2), expected: "any of (value greater than 2 or value equal to 2)"},
		{description: "LessThan", actual: 2, matcher: is.LessThan(1), expected: "value less than 1"},
		{description: "LessThanOrEqualTo", actual: 2, matcher: is.LessThanOrEqualTo(1), expected: "any of (value less than 1 or value equal to 1)"},
		{description: "Not", actual: 2, matcher: is.Not(is.EqualTo(2)), expected: "\nExpected: not(value equal to 2)\n     but: 2\n"},
		{description: "Nil", actual: 1, matcher: is.Nil(), expected: "value equal to <nil>"},
		{description: "Contains", actual: []string{"Foo", "Bar"}, matcher: is.ValueContaining([]string{"Baz", "Bing"}), expected: "something that contains [Baz Bing]"},
		{description: "MatchesPattern", actual: "blarney stone", matcher: is.MatchForPattern("~123.?.*"), expected: "a value that matches pattern ~123.?.*"},
		{description: "MatchesPattern (invalid regex)", actual: "blarney stone", matcher: is.MatchForPattern("+++"), expected: "error parsing regexp: missing argument to repetition operator: `+`"},
		{description: "Prefix", actual: "blarney stone", matcher: has.Prefix("123"), expected: "value with prefix 123"},
		{description: "AllOf", actual: "abc", matcher: is.AllOf(is.EqualTo("abc"), is.ValueContaining("e")), expected: "all of (value equal to abc and something that contains e)"},
		{description: "AnyOf", actual: "abc", matcher: is.AnyOf(is.EqualTo("efg"), is.ValueContaining("e")), expected: "any of (value equal to efg or something that contains e)"},
		{description: "HasKey", actual: map[string]bool{"hi": true}, matcher: has.Key("foo"), expected: "map has key 'foo'"},
		{description: "HasKeys", actual: map[string]bool{"hi": true, "bye": false}, matcher: has.AllKeys("hi", "foo"), expected: "map has keys '[hi foo]'"},
		{description: "LengthOf Composed", actual: "a", matcher: has.Length(is.GreaterThan(2)), expected: "value with length value greater than 2"},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, test.matcher)
		if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
			testing.Errorf("%s did not fail with expected desc, got: %s", test.description, stubTestingT.MockTestOutput)
		}
	}
}

func TestHasFieldDescribesMismatch(testing *testing.T) {
	type T struct {
		F string
		B string
	}
	expected := "X"
	then.AssertThat(stubTestingT, new(T), has.FieldNamed(expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "struct with field X") &&
		!strings.Contains(stubTestingT.MockTestOutput, "T{F string B string}") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}

func TestHasFunctionDescribesMismatch(testing *testing.T) {
	type MyType interface {
		F() string
		B() string
	}
	actual := new(MyType)
	expected := "X"
	then.AssertThat(stubTestingT, actual, has.FunctionNamed(expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "interface with function X") &&
		!strings.Contains(stubTestingT.MockTestOutput, "MyType{B()F()}") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}
