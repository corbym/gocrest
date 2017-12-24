package gocrest_test

import (
	"testing"
	"fmt"
	"gocrest"
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
