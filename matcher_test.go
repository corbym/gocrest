package gocrest_test

import (
	"testing"
	"fmt"
	"gocrest"
	"strings"
)

const failedTest = true
const passedTest = false

type MockTestingT struct {
	testStatus     bool
	mockTestOutput string
}

func (t *MockTestingT) Logf(format string, args ...interface{}) {
	t.mockTestOutput = fmt.Sprintf(format, args...)
	t.testStatus = failedTest
}

func (t *MockTestingT) Errorf(format string, args ...interface{}) {
	t.mockTestOutput = fmt.Sprintf(format, args...)
	t.testStatus = failedTest
}

func (t *MockTestingT) FailNow() {
	t.testStatus = failedTest
}

var mockTestingT *MockTestingT

func init() {
	mockTestingT = new(MockTestingT)
}

func TestAssertThatTwoValuesAreEqualOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: passedTest},
		{actual: 1.12, expected: 1.12, shouldFail: passedTest},
		{actual: 1, expected: 2, shouldFail: failedTest},
		{actual: "hi", expected: "bees", shouldFail: failedTest},
	}
	for _, test := range equalsItems {
		gocrest.AssertThat(mockTestingT, test.actual, gocrest.EqualTo(test.expected))
		if mockTestingT.testStatus != test.shouldFail {
			testing.Errorf("assertThat(%v, EqualTo(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, mockTestingT.testStatus)
		}
	}
}

func TestNotReturnsTheOppositeOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.Not(gocrest.EqualTo(2)))
	if mockTestingT.testStatus == passedTest {
		testing.Error("Not(EqualTo) did not fail the test")
	}
}

func TestNotReturnsTheNotDescriptionOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 2, gocrest.Not(gocrest.EqualTo(2)))
	if mockTestingT.mockTestOutput != "expected: not(value equal to 2) but was: 2" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.mockTestOutput)
	}
}

func TestAssertThatFailsTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2))
	if mockTestingT.testStatus == passedTest {
		testing.Error("1 EqualTo 2 did not fail test")
	}
}

func TestEqualToFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2))
	if mockTestingT.mockTestOutput != "expected: value equal to 2 but was: 1" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.mockTestOutput)
	}
}

func TestIsNilMatches(testing *testing.T) {
	gocrest.AssertThat(testing, nil, gocrest.IsNil())
}

func TestIsNilFails(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 2, gocrest.IsNil())
	if mockTestingT.testStatus == passedTest {
		testing.Fail()
	}
}

func TestIsNilHasDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.IsNil())
	if mockTestingT.mockTestOutput != "expected: value equal to <nil> but was: 1" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.mockTestOutput)
	}
}

func TestContainsDescriptionTest(testing *testing.T) {
	list := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(mockTestingT, list, gocrest.Contains(expectedList))
	if mockTestingT.mockTestOutput != "expected: something that contains [Baz Bing] but was: [Foo Bar]" {
		testing.Errorf("did not get expected description, got %s", mockTestingT.mockTestOutput)
	}
}

func TestContainsFailsForTwoStringArraysTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expectedList))
	if mockTestingT.testStatus == passedTest {
		testing.Fail()
	}
}

func TestContainsFailsForTwoIntArraysTest(testing *testing.T) {
	actualList := []int{12, 13}
	expectedList := []int{14, 15}
	gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expectedList))
	if mockTestingT.testStatus == passedTest {
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
	if mockTestingT.testStatus == passedTest {
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
	if mockTestingT.testStatus == passedTest {
		testing.Error("did not fail test")
	}
}

func TestMatchesPatternDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "~123.?.*"
	gocrest.AssertThat(mockTestingT, actual, gocrest.MatchesPattern(expected))
	if mockTestingT.mockTestOutput != "expected: a value that matches pattern ~123.?.* but was: blarney stone" {
		testing.Errorf("incorrect description: %s", mockTestingT.mockTestOutput)
	}
}

func TestMatchesPatternWithErrorDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "+++"
	gocrest.AssertThat(mockTestingT, actual, gocrest.MatchesPattern(expected))
	if !strings.Contains(mockTestingT.mockTestOutput, "error parsing regexp: missing argument to repetition operator: `+`") {
		testing.Errorf("incorrect description: %s", mockTestingT.mockTestOutput)
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
	if mockTestingT.testStatus == passedTest {
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
	if mockTestingT.mockTestOutput != "expected: interface with function X but was: MyType{B()F()}" {
		testing.Errorf("incorrect description:%s", mockTestingT.mockTestOutput)
	}
}
