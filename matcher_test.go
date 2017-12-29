package gocrest_test

import (
	"testing"
	"strings"
	"gocrest/then"
	"gocrest/is"
	"gocrest/has"
)

var stubTestingT *StubTestingT

func init() {
	stubTestingT = new(StubTestingT)
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

func TestAssertThatEqualToFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.EqualTo(2).Reason("arithmetic is wrong"))
	if stubTestingT.MockTestOutput != "arithmetic is wrong\nExpected: value equal to 2\n     but: 1" {
		testing.Errorf("did not get expected description, got: %s", stubTestingT.MockTestOutput)
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

func TestEmptyFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, map[string]bool{"foo": true}, is.Empty())
	if !strings.Contains(stubTestingT.MockTestOutput, "empty value") {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
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
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: true}, // cannot compare complex types, so fails
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.GreaterThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, GreaterThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestGreaterThanFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.GreaterThan(2))
	if !strings.Contains(stubTestingT.MockTestOutput, "value greater than 2") {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
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

func TestLessThanFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, 2, is.LessThan(1))
	if stubTestingT.MockTestOutput != "\nExpected: value less than 1\n     but: 2" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestNotReturnsTheOppositeOfGivenMatcher(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.Not(is.EqualTo(2)))
	if !stubTestingT.HasFailed() {
		testing.Error("Not(EqualTo) did not fail the test")
	}
}

func TestNotReturnsTheNotDescriptionOfGivenMatcher(testing *testing.T) {
	then.AssertThat(stubTestingT, 2, is.Not(is.EqualTo(2)))
	if stubTestingT.MockTestOutput != "\nExpected: not(value equal to 2)\n     but: 2" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
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

func TestIsNilHasDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.Nil())
	if !strings.Contains(stubTestingT.MockTestOutput, "value equal to <nil>") {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestContainsDescriptionTest(testing *testing.T) {
	list := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	then.AssertThat(stubTestingT, list, is.ValueContaining(expectedList))
	if !strings.Contains(stubTestingT.MockTestOutput, "something that contains [Baz Bing]") {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
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

func TestMatchesPatternDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "~123.?.*"
	then.AssertThat(stubTestingT, actual, is.MatchForPattern(expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "a value that matches pattern ~123.?.*") {
		testing.Errorf("incorrect description: %s", stubTestingT.MockTestOutput)
	}
}

func TestMatchesPatternWithErrorDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "+++"
	then.AssertThat(stubTestingT, actual, is.MatchForPattern(expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "error parsing regexp: missing argument to repetition operator: `+`") {
		testing.Errorf("incorrect description: %s", stubTestingT.MockTestOutput)
	}
}

func TestHasFunctionPasses(testing *testing.T) {
	type MyType interface {
		N() int
		F() string
	}
	actual := new(MyType)
	expected := "F"
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

func TestAllOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AllOf(is.EqualTo("abc"), is.ValueContaining("e")))
	if !strings.Contains(stubTestingT.MockTestOutput, "all of (value equal to abc and something that contains e)") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
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

func TestAnyOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AnyOf(is.EqualTo("efg"), is.ValueContaining("e")))
	if !strings.Contains(stubTestingT.MockTestOutput, "any of (value equal to efg or something that contains e") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}
