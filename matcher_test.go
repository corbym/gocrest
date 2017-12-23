package gocrest_test

import (
	"testing"
	"strconv"
	"fmt"
	"gocrest"
)

var output string

type MockTestingT struct{}

func (MockTestingT) Logf(format string, args ...interface{}) {
	output = fmt.Sprintf(format, args...)
}

func (MockTestingT) Errorf(format string, args ...interface{}) {
	output = fmt.Sprintf(format, args...)
}
func (MockTestingT) FailNow() {}

var mockTestingT = new(MockTestingT)

func TestAssertThatTwoValuesAreEqualOrNot(testing *testing.T) {

	var equalsItems = []struct {
		actual   interface{} // input
		expected interface{} // expected
		result   bool        // result if equal
	}{
		{actual: 1, expected: 1, result: true},
		{actual: 1, expected: 2, result: false},
		{actual: "hi", expected: "hi", result: true},
		{actual: 1.12, expected: 1.12, result: true},
	}
	for _, test := range equalsItems {
		wasEqual := gocrest.AssertThat(mockTestingT, test.actual, gocrest.EqualTo(test.expected))
		if wasEqual != test.result {
			testing.Error("wanted " + strconv.FormatBool(test.result) + ", got " + strconv.FormatBool(wasEqual))
		}
	}
}

func TestAssertThatFailsTest(testing *testing.T) {
	if gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2)) != false {
		testing.FailNow()
	}
}

func TestEqualToFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.EqualTo(2))
	if output != "expected: value equal to 2 but was: 1" {
		testing.Error("did not get expected description, got " + output)
	}
}

func TestIsNilMatches(testing *testing.T) {
	gocrest.AssertThat(testing, nil, gocrest.IsNil())
}

func TestIsNilDoesNotMatch(testing *testing.T) {
	result := gocrest.AssertThat(mockTestingT, 2, gocrest.IsNil())
	if result {
		testing.Fail()
	}
}

func TestIsNilHasDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(mockTestingT, 1, gocrest.IsNil())
	if output != "expected: value equal to <nil> but was: 1" {
		testing.Error("did not get expected description, got " + output)
	}
}

func TestContainsFailsWithDescriptionTest(testing *testing.T) {
	list := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(mockTestingT, list, gocrest.Contains(expectedList))
	if output != "expected: something that contains [Baz Bing] but was: [Foo Bar]" {
		testing.Error("did not get expected description, got " + output)
	}
}

func TestContainsFailsForTwoArraysWithDescriptionTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	result := gocrest.AssertThat(mockTestingT, actualList, gocrest.Contains(expectedList))
	if result {
		testing.Fail()
	}
}

func TestContainsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Foo"
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestContainsForMap(testing *testing.T) {
	actualList := make(map[string]string)
	actualList["bing"] = "boop"
	actualList["bling"] = "bling"
	expected := make(map[string]string)
	expected["bing"] = "boop"

	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}
