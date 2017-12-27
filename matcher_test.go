package gocrest_test

import (
	"testing"
	"gocrest"
	"strings"
)

var mockTestingT *gocrest.MockTestingT

func init() {
	mockTestingT = new(gocrest.MockTestingT)
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
	}
	for _, test := range equalsItems {
		gocrest.AssertThat(mockTestingT, test.actual, gocrest.EqualTo(test.expected))
		if mockTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, EqualTo(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, mockTestingT.HasFailed())
		}
	}
}

func TestNotReturnsTheOppositeOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.Not(gocrest.EqualTo(2)))
	if !mockTestingT.HasFailed() {
		testing.Error("Not(EqualTo) did not fail the test")
	}
}

func TestNotReturnsTheNotDescriptionOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 2, gocrest.Not(gocrest.EqualTo(2)))
	if mockTestingT.MockTestOutput != "expected: not(value equal to 2) but was: 2" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.MockTestOutput)
	}
}

func TestAssertThatFailsTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2))
	if !mockTestingT.HasFailed() {
		testing.Error("1 EqualTo 2 did not fail test")
	}
}

func TestEqualToFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2))
	if mockTestingT.MockTestOutput != "expected: value equal to 2 but was: 1" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.MockTestOutput)
	}
}

func TestIsNilMatches(testing *testing.T) {
	gocrest.AssertThat(testing, nil, gocrest.IsNil())
}

func TestIsNilFails(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 2, gocrest.IsNil())
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestIsNilHasDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.IsNil())
	if mockTestingT.MockTestOutput != "expected: value equal to <nil> but was: 1" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.MockTestOutput)
	}
}

func TestContainsDescriptionTest(testing *testing.T) {
	list := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(mockTestingT, list, gocrest.Contains(expectedList))
	if mockTestingT.MockTestOutput != "expected: something that contains [Baz Bing] but was: [Foo Bar]" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.MockTestOutput)
	}
}

func TestContainsFailsForTwoStringArraysTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expectedList))
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoIntArraysTest(testing *testing.T) {
	actualList := []int{12, 13}
	expectedList := []int{14, 15}
	gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expectedList))
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Foo"
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestContainsFailsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Moo"
	gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expected))
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForSlice(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Baz", "Bing", "Bong"}
	gocrest.AssertThat(testing, actualList[2:2], gocrest.Contains(expected[2:2]))
}

func TestContainsForList(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Boom", "Bong", "Bar"}
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestMapContainsMap(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	expected := map[string]string{
		"bing": "boop",
	}

	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestStringContainsString(testing *testing.T) {
	actualList := "abcd"
	expected := "bc"
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestMatchesPatternMatchesString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^blarney.*"
	gocrest.AssertThat(testing, actual, gocrest.MatchesPattern(expected))
}

func TestMatchesPatternDoesNotMatchString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^123.?.*"
	gocrest.AssertThat(mockTestingT, actual, gocrest.MatchesPattern(expected))
	if !mockTestingT.HasFailed() {
		testing.Error("did not fail test")
	}
}

func TestMatchesPatternDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "~123.?.*"
	gocrest.AssertThat(mockTestingT, actual, gocrest.MatchesPattern(expected))
	if mockTestingT.MockTestOutput != "expected: a value that matches pattern ~123.?.* but was: blarney stone" {
		testing.Errorf("incorrect description: %s", mockTestingT.MockTestOutput)
	}
}

func TestMatchesPatternWithErrorDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "+++"
	gocrest.AssertThat(mockTestingT, actual, gocrest.MatchesPattern(expected))
	if !strings.Contains(mockTestingT.MockTestOutput, "error parsing regexp: missing argument to repetition operator: `+`") {
		testing.Errorf("incorrect description: %s", mockTestingT.MockTestOutput)
	}
}

func TestHasFunctionPasses(testing *testing.T) {
	type MyType interface {
		N() int
		F() string
	}
	actual := new(MyType)
	expected := "F"
	gocrest.AssertThat(testing, actual, gocrest.HasFunctionNamed(expected))
}

func TestHasFunctionDoesNotPass(testing *testing.T) {
	type MyType interface {
		F() string
	}
	actual := new(MyType)
	expected := "E"
	gocrest.AssertThat(mockTestingT, actual, gocrest.HasFunctionNamed(expected))
	if !mockTestingT.HasFailed() {
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
	gocrest.AssertThat(mockTestingT, actual, gocrest.HasFunctionNamed(expected))
	if mockTestingT.MockTestOutput != "expected: interface with function X but was: MyType{B()F()}" {
		testing.Errorf("incorrect description:%s", mockTestingT.MockTestOutput)
	}
}

func TestAllOfMatches(testing *testing.T) {
	actual := "abcdef"
	gocrest.AssertThat(testing, actual, gocrest.AllOf(gocrest.EqualTo("abcdef"), gocrest.Contains("e")))
}

func TestAllOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(mockTestingT, actual, gocrest.AllOf(gocrest.EqualTo("abc"), gocrest.Contains("e")))
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAllOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(mockTestingT, actual, gocrest.AllOf(gocrest.EqualTo("abc"), gocrest.Contains("e")))
	if mockTestingT.MockTestOutput != "expected: all of (value equal to abc and something that contains e) but was: abc" {
		testing.Errorf("incorrect description:%s", mockTestingT.MockTestOutput)
	}
}

func TestAnyOfMatches(testing *testing.T) {
	actual := "abcdef"
	gocrest.AssertThat(testing, actual, gocrest.AnyOf(gocrest.EqualTo("abcdef"), gocrest.Contains("g")))
}

func TestAnyOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(mockTestingT, actual, gocrest.AnyOf(gocrest.EqualTo("efg"), gocrest.Contains("e")))
	if !mockTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAnyOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(mockTestingT, actual, gocrest.AnyOf(gocrest.EqualTo("efg"), gocrest.Contains("e")))
	if mockTestingT.MockTestOutput != "expected: any of (value equal to efg or something that contains e) but was: abc" {
		testing.Errorf("incorrect description:%s", mockTestingT.MockTestOutput)
	}
}
