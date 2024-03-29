package gocrest_test

import (
	"bytes"
	"fmt"
	"github.com/corbym/gocrest"
	"github.com/corbym/gocrest/by"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"io"
	"strings"
	"testing"
	"time"
)

var stubTestingT *StubTestingT

func init() {
	stubTestingT = new(StubTestingT)
}

func TestEqualTo(t *testing.T) {
	then.AssertThat(t, "foo", is.EqualTo("foo"))
	then.AssertThat(t, 1, is.EqualTo(1))
	then.AssertThat(t, []string{"a"}, is.EqualTo([]string{"a"}))
}

func TestMixedMatcher(t *testing.T) {
	then.AssertThat(t, "abcdef", is.AllOf(is.StringContaining("abc"), is.LessThan("ghi")))
}

func TestNil(testing *testing.T) {
	values := struct {
		actual *string
	}{
		actual: nil,
	}
	then.AssertThat(testing, values.actual, is.NilPtr[string]())
}

func TestError(testing *testing.T) {
	values := struct {
		actual error
	}{
		actual: nil,
	}
	then.AssertThat(testing, values.actual, is.Nil())
}

func TestHasLengthStringMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     string
		expected   int
		shouldFail bool
	}{
		{actual: "", expected: 0, shouldFail: false},
		{actual: "a", expected: 1, shouldFail: false},
		{actual: "a", expected: 1, shouldFail: false},
		{actual: "1", expected: 1, shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(testing, "a", has.StringLength(1))
		then.AssertThat(stubTestingT, test.actual, has.StringLength(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestMapHasLengthOne(t *testing.T) {
	then.AssertThat(t, map[string]bool{"hello": true}, has.MapLength[string, bool](1))
}
func TestHasLengthMapMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     map[string]bool
		expected   *gocrest.Matcher[int]
		shouldFail bool
	}{
		{actual: map[string]bool{"helloa": true}, expected: is.LessThan(1), shouldFail: true},
		{actual: map[string]bool{"hellob": true}, expected: is.LessThanOrEqualTo(2), shouldFail: false},
		{actual: map[string]bool{"helloc": true}, expected: is.GreaterThan(2), shouldFail: true},
		{actual: map[string]bool{"hellod": true}, expected: is.GreaterThanOrEqualTo(1), shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.MapLengthMatching[string, bool](test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestHasLengthArrayMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     []int
		expected   int
		shouldFail bool
	}{
		{actual: []int{}, expected: 0, shouldFail: false},
		{actual: []int{1}, expected: 1, shouldFail: false},
		{actual: []int{1}, expected: 2, shouldFail: true},
		{actual: []int{1, 2}, expected: 2, shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.Length[int](test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestHasLengthMatchesArrayMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     []int
		expected   int
		shouldFail bool
	}{
		{actual: []int{}, expected: 0, shouldFail: false},
		{actual: []int{1}, expected: 1, shouldFail: false},
		{actual: []int{1}, expected: 2, shouldFail: true},
		{actual: []int{1, 2}, expected: 2, shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.LengthMatching[int](is.EqualTo(test.expected)))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestHasLengthArrayStringMatchesOrNot(testing *testing.T) {
	var hasLengthItems = []struct {
		actual     []string
		expected   int
		shouldFail bool
	}{
		{actual: nil, expected: 2, shouldFail: true},
		{actual: []string{}, expected: 0, shouldFail: false},
		{actual: []string{"foo"}, expected: 1, shouldFail: false},
		{actual: []string{"foo"}, expected: 2, shouldFail: true},
		{actual: []string{"foo", "bar"}, expected: 2, shouldFail: false},
	}
	for _, test := range hasLengthItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.Length[string](test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, has.Length(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoValuesAreEqualOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     any
		expected   any
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
		actual     string
		shouldFail bool
	}{
		{actual: "hi", shouldFail: true},
		{actual: "", shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.EmptyString())
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, Empty()) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestEmptyMapIsEmptyPasses(testing *testing.T) {
	var equalsItems = []struct {
		actual     map[string]bool
		shouldFail bool
	}{
		{actual: map[string]bool{"hello": true}, shouldFail: true},
		{actual: map[string]bool{}, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.EmptyMap[string, bool]())
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, EmptyMap()) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestEmptyArrayIsEmptyPasses(testing *testing.T) {
	var equalsItems = []struct {
		actual     []string
		shouldFail bool
	}{
		{actual: []string{}, shouldFail: false},
		{actual: []string{"boo"}, shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.Empty[string]())
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, Empty()) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoIntValuesAreGreaterThanOrNotFails(testing *testing.T) {
	then.AssertThat(stubTestingT, 1, is.GreaterThan(1))
	hasFailed(testing)
	then.AssertThat(stubTestingT, 1.12, is.GreaterThan(1.12))
	hasFailed(testing)
}

func hasFailed(testing *testing.T) {
	if !stubTestingT.HasFailed() {
		testing.Errorf(testing.Name())
	}
}
func TestAssertThatTwoValuesAreGreaterThanOrNotFails(testing *testing.T) {
	var equalsItems = []struct {
		actual     string
		expected   string
		shouldFail bool
	}{
		{actual: "zzz", expected: "aaa", shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.GreaterThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, GreaterThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoFloatValuesAreGreaterThanOrNot(testing *testing.T) {
	then.AssertThat(testing, float32(1.12), is.GreaterThan(float32(1.0)))
	then.AssertThat(testing, 2, is.GreaterThan(1))
	then.AssertThat(testing, 1.24, is.GreaterThan(1.0))
}
func TestAssertThatTwoIntsValuesAreGreaterThanOrNot(testing *testing.T) {
	then.AssertThat(testing, uint(3), is.GreaterThan(uint(1)))
	then.AssertThat(testing, uint16(4), is.GreaterThan(uint16(1)))
	then.AssertThat(testing, uint32(6), is.GreaterThan(uint32(1)))
	then.AssertThat(testing, uint64(7), is.GreaterThan(uint64(1)))
	then.AssertThat(testing, uint64(8), is.GreaterThan(uint64(1)))
	then.AssertThat(testing, int16(9), is.GreaterThan(int16(1)))
	then.AssertThat(testing, int32(10), is.GreaterThan(int32(1)))
	then.AssertThat(testing, int64(11), is.GreaterThan(int64(1)))
	then.AssertThat(testing, int64(12), is.GreaterThan(int64(1)))
}
func TestAssertThatHasLengthFailsWithDescriptionTest(testing *testing.T) {
	then.AssertThat(stubTestingT, "a", has.StringLength(2))
	if !strings.Contains(stubTestingT.MockTestOutput, "value with length 2") {
		testing.Errorf("did not get expected description, got: %s", stubTestingT.MockTestOutput)
	}
}

func TestAssertThatTwoIntsAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     int
		expected   int
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 1, expected: 2, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoFloat32ValuesAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     float32
		expected   float32
		shouldFail bool
	}{
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: float32(1.0), expected: float32(1.12), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoFloat64sAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     float64
		expected   float64
		shouldFail bool
	}{
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: 1.0, expected: 1.24, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoUintsAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     uint16
		expected   uint16
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: uint16(1), expected: uint16(4), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoUints64sAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     uint64
		expected   uint64
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: uint64(1), expected: uint64(7), shouldFail: false},
		{actual: uint64(1), expected: uint64(8), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestAssertThatTwoIntValuesAreLessThanOrEqualToPassesOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     int16
		expected   int16
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: int16(1), expected: int16(9), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThanOrEqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoInt32ValuesAreLessThanOrEqualToPassesOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     int32
		expected   int32
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: int32(1), expected: int32(10), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThanOrEqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoInt64ValuesAreLessThanOrEqualToPassesOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     int64
		expected   int64
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: int64(1), expected: int64(11), shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT = new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.LessThanOrEqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}
func TestAssertThatTwoStringValuesAreLessThanOrEqualToPassesOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     string
		expected   string
		shouldFail bool
	}{
		{actual: "aaa", expected: "zzz", shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: true},
		{actual: "aaa", expected: "aaa", shouldFail: false},
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

func TestNotReturnsTheSubMatcherActual(testing *testing.T) {
	not := is.Not(has.StringLength(1))
	not.Matches("a")
	then.AssertThat(testing, not.Actual,
		is.EqualTo("length was 1"))
}

func TestAnyofReturnsTheSubMatcherActual(testing *testing.T) {
	anyOf := is.AnyOf(has.StringLength(1), is.EqualTo("a"))
	anyOf.Matches("a")
	then.AssertThat(testing, anyOf.Actual,
		is.EqualTo("actual <a> length was 1"))
}

func TestAllofReturnsTheSubMatcherActual(testing *testing.T) {
	anyOf := is.AllOf(has.StringLength(1), is.EqualTo("a"))
	anyOf.Matches("a")
	then.AssertThat(testing, anyOf.Actual,
		is.EqualTo("actual <a> length was 1"))
}

func TestIsNilMatches(testing *testing.T) {
	then.AssertThat(testing, nil, is.Nil())
}

func TestIsNilFails(testing *testing.T) {
	var actual = 2
	then.AssertThat(stubTestingT, &actual, is.NilPtr[int]())
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}
func TestIsNilArrayFails(testing *testing.T) {
	var actual = []int{1, 2}
	then.AssertThat(stubTestingT, actual, is.NilArray[int]())
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}
func TestIsNilMapFails(testing *testing.T) {
	var actual = map[string]string{"a": "b"}
	then.AssertThat(stubTestingT, actual, is.NilMap[string, string]())
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoStringArraysTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	then.AssertThat(stubTestingT, actualList, is.ArrayContaining(expectedList...))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoIntArraysTest(testing *testing.T) {
	actualList := []int{12, 13}
	expectedList := []int{14, 15}
	then.AssertThat(stubTestingT, actualList, is.ArrayContaining(expectedList...))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Foo"
	then.AssertThat(testing, actualList, is.ArrayContaining(expected))
}

func TestContainsFailsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Moo"
	then.AssertThat(stubTestingT, actualList, is.ArrayContaining(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForSlice(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Baz", "Bing", "Bong"}
	then.AssertThat(testing, actualList[2:2], is.ArrayContaining(expected[2:2]...))
}

func TestContainsForList(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Boom", "Bong", "Bar"}
	then.AssertThat(testing, actualList, is.ArrayContaining(expected...))
}

func TestContainsForVariadic(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	then.AssertThat(testing, actualList, is.ArrayContaining("Boom", "Bong", "Bar"))
}

func TestContainsForVariadicMatchers(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	then.AssertThat(testing, actualList, is.ArrayMatching(is.EqualTo("Boom"), has.Suffix("ng"), has.Prefix("Ba")))
}

func TestMapContainsMap(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	expected := map[string]string{
		"bing": "boop",
	}

	then.AssertThat(testing, actualList, is.MapContaining(expected))
}

func TestMapContainsValues(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	then.AssertThat(testing, actualList, is.MapContainingValues[string]("boop", "bling"))
}

func TestMapMatchesValues(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	then.AssertThat(testing, actualList, is.MapMatchingValues[string](is.EqualTo("boop"), is.EqualTo("bling")))
}

func TestStringContains_String(testing *testing.T) {
	actualList := "abcd"
	expected := "bc"
	then.AssertThat(testing, actualList, is.StringContaining(expected))
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
	then.AssertThat(testing, actual, has.FunctionNamed[*MyType](expected))
}

func TestHasFunctionDoesNotPass(testing *testing.T) {
	type MyType interface {
		F() string
	}
	actual := new(MyType)
	expected := "E"
	then.AssertThat(stubTestingT, actual, has.FunctionNamed[*MyType](expected))
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
	then.AssertThat(testing, actual, has.FieldNamed[*T](expected))
}

func TestHasFieldDoesNotPass(testing *testing.T) {
	type T struct {
		F int
	}
	actual := new(T)
	expected := "E"
	then.AssertThat(stubTestingT, actual, has.FieldNamed[*T](expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAllOfMatches(testing *testing.T) {
	actual := "abcdef"
	then.AssertThat(testing, actual, is.AllOf(is.EqualTo("abcdef"), is.StringContaining("e")))
}

func TestAllOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AllOf(is.EqualTo("abc"), is.StringContaining("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAnyOfMatches(testing *testing.T) {
	actual := "abcdef"
	then.AssertThat(testing, actual, is.AnyOf(is.EqualTo("abcdef"), is.StringContaining("g")))
}

func TestAnyOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	then.AssertThat(stubTestingT, actual, is.AnyOf(is.EqualTo("efg"), is.StringContaining("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestHasKeyMatches(testing *testing.T) {
	type T struct{}
	expectedT := new(T)
	var equalsItems = []struct {
		actual     map[*T]bool
		expected   *T
		shouldFail bool
	}{
		{actual: map[*T]bool{expectedT: true}, expected: expectedT, shouldFail: false},
	}
	for _, test := range equalsItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.Key[*T, bool](test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("unexpected result HasKey: wanted fail was %v but failed %v", test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestHasKeysMatches(testing *testing.T) {
	then.AssertThat(testing, map[string]bool{"hi": true, "bye": true}, has.AllKeys[string, bool]("hi", "bye"))
	type T struct{}
	expectedT := new(T)
	secondExpectedT := new(T)
	then.AssertThat(testing, map[*T]bool{expectedT: true, secondExpectedT: true}, has.AllKeys[*T, bool](expectedT, secondExpectedT))
}

func TestBoolMatcherDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      bool
		matcher     *gocrest.Matcher[bool]
		expected    string
	}{
		{description: "is true", actual: false, matcher: is.True(), expected: "is true"},
		{description: "is false", actual: true, matcher: is.False(), expected: "is false"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}
func TestTypeNameMatcherDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      *testing.T
		matcher     *gocrest.Matcher[*testing.T]
		expected    string
	}{
		{description: "has type T", actual: t, matcher: has.TypeName[*testing.T]("string"), expected: "has type <string>"},
		{description: "has type T matcher", actual: t, matcher: has.TypeNameMatches[*testing.T](is.EqualTo("string")), expected: "has type value equal to <string>"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}
func TestSizeMatcherDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      int
		matcher     *gocrest.Matcher[int]
		expected    string
	}{
		{description: "EqualTo.Reasonf", actual: 1, matcher: is.EqualTo(2).Reasonf("arithmetic %s is wrong", "foo"), expected: "arithmetic foo is wrong"},
		{description: "EqualTo.Reason", actual: 1, matcher: is.EqualTo(2).Reason("arithmetic is wrong"), expected: "arithmetic is wrong\nExpected: value equal to <2>\n     but: <1>\n"},
		{description: "Not", actual: 2, matcher: is.Not(is.EqualTo(2)), expected: "\nExpected: not(value equal to <2>)\n     but: <2>\n"},
		{description: "GreaterThan", actual: 1, matcher: is.GreaterThan(2), expected: "value greater than <2>"},
		{description: "GreaterThanOrEqual", actual: 1, matcher: is.GreaterThanOrEqualTo(2), expected: "any of (value greater than <2> or value equal to <2>)"},
		{description: "LessThan", actual: 2, matcher: is.LessThan(1), expected: "value less than <1>"},
		{description: "LessThanOrEqualTo", actual: 2, matcher: is.LessThanOrEqualTo(1), expected: "any of (value less than <1> or value equal to <1>)"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}
func TestSizeMapMatcherDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      map[string]bool
		matcher     *gocrest.Matcher[map[string]bool]
		expected    string
	}{
		{description: "Empty", actual: map[string]bool{"foo": true}, matcher: is.EmptyMap[string, bool](), expected: "empty value"},
		{description: "HasKey", actual: map[string]bool{"hi": true}, matcher: has.Key[string, bool]("foo"), expected: "map has key 'foo'"},
		{description: "HasKeys", actual: map[string]bool{"hi": true, "bye": false}, matcher: has.AllKeys[string, bool]("hi", "foo"), expected: "map has keys '[hi foo]'"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}
func TestArrayContainsMatcherDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      []string
		matcher     *gocrest.Matcher[[]string]
		expected    string
	}{
		{description: "ArrayContaining", actual: []string{"Foo", "Bar"}, matcher: is.ArrayContaining("Baz", "Bing"), expected: "something that contains <Baz> and <Bing>"},
		{description: "ValueContainArrayMatching", actual: []string{"Foo", "Bar"}, matcher: is.ArrayMatching(is.EqualTo("Baz"), is.EqualTo("Bing")), expected: "something that contains <value equal to <Baz>> and <value equal to <Bing>>"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}
func TestStringMatchersDescription(t *testing.T) {
	var equalsItems = []struct {
		description string
		actual      string
		matcher     *gocrest.Matcher[string]
		expected    string
	}{
		{description: "MatchesPattern", actual: "blarney stone", matcher: is.MatchForPattern("~123.?.*"), expected: "a value that matches pattern ~123.?.*"},
		{description: "MatchesPattern (invalid regex)", actual: "blarney stone", matcher: is.MatchForPattern("+++"), expected: "error parsing regexp: missing argument to repetition operator: `+`"},
		{description: "Prefix", actual: "blarney stone", matcher: has.Prefix("123"), expected: "value with prefix 123"},
		{description: "AllOf", actual: "abc", matcher: is.AllOf(is.EqualTo("abc"), is.StringContaining("e", "f")), expected: "something that contains [e f]"},
		{description: "AnyOf", actual: "abc", matcher: is.AnyOf(is.EqualTo("efg"), is.StringContaining("e")), expected: "any of (value equal to <efg> or something that contains [e])"},
		{description: "LengthOf Composed", actual: "a", matcher: has.StringLengthMatching(is.GreaterThan(2)), expected: "value with length value greater than <2>"},
		{description: "EqualToIgnoringWhitespace", actual: "a b c", matcher: is.EqualToIgnoringWhitespace("b c d"), expected: "ignoring whitespace value equal to <b c d>"},
	}
	for _, test := range equalsItems {
		t.Run(test.description, func(innerT *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, test.actual, test.matcher)
			if !strings.Contains(stubTestingT.MockTestOutput, test.expected) {
				innerT.Errorf("%s did not fail with expected desc <%s>, got: %s", test.description, test.expected, stubTestingT.MockTestOutput)
			}
		})
	}
}

func TestAllOfDescribesOnlyMismatches(testing *testing.T) {
	stubTestingT := new(StubTestingT)
	then.AssertThat(stubTestingT, "abc", is.AllOf(
		is.EqualTo("abc"),
		is.StringContaining("e", "f"),
		is.EmptyString(),
	))
	if !strings.Contains(stubTestingT.MockTestOutput, "Expected: something that contains [e f] and empty value\n") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}

func TestHasFieldDescribesMismatch(testing *testing.T) {
	type T struct {
		F string
		B string
	}
	expected := "X"
	then.AssertThat(stubTestingT, new(T), has.FieldNamed[*T](expected))
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
	then.AssertThat(stubTestingT, actual, has.FunctionNamed[*MyType](expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "interface with function X") &&
		!strings.Contains(stubTestingT.MockTestOutput, "MyType{B()F()}") {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}

func TestEqualToIgnoringWhitespace(t *testing.T) {
	var ignoreWhitespaceItems = []struct {
		actual     string
		expected   string
		shouldFail bool
	}{
		{actual: "a bc", expected: "a bc", shouldFail: false},
		{actual: "a b c", expected: "a bc", shouldFail: false},
		{actual: "abc", expected: "a", shouldFail: true},
		{actual: "abc \n", expected: "a", shouldFail: true},
		{actual: "%^&*abc \n\t\t", expected: "%^&*abc\n", shouldFail: false},
	}
	for _, test := range ignoreWhitespaceItems {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, is.EqualToIgnoringWhitespace(test.expected))
		then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reasonf("<%s>, should Fail(%v) for <%s>", test.actual, test.shouldFail, test.expected))
	}
}

type private struct {
}
type Foo struct {
}

func TestTypeName(t *testing.T) {
	pri := new(private)
	priAnd := &private{}
	pubAnd := &Foo{}
	pub := new(Foo)

	tests := []struct {
		actual     interface{}
		expected   string
		shouldFail bool
	}{
		{t, "*testing.T", false},
		{t, "foob", true},
		{pri, "*gocrest_test.private", false},
		{priAnd, "*gocrest_test.private", false},
		{private{}, "gocrest_test.private", false},
		{pub, "*gocrest_test.Foo", false},
		{pubAnd, "*gocrest_test.Foo", false},
		{&pubAnd, "**gocrest_test.Foo", false},
		{Foo{}, "gocrest_test.Foo", false},
		{"foo", "string", false},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			stubTestingT := new(StubTestingT)
			then.AssertThat(stubTestingT, tt.actual, has.TypeName[any](tt.expected))

			then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(tt.shouldFail).Reason(stubTestingT.MockTestOutput))
		})
	}
}

func TestNilError(t *testing.T) {
	actual := nilResponse()

	then.AssertThat(t, actual, is.Nil())
}

func nilResponse() error {
	return nil
}

func TestEveryStringElement(t *testing.T) {
	tests := []struct {
		actual     []string
		expected   []*gocrest.Matcher[string]
		shouldFail bool
	}{
		{
			actual:     []string{"test1", "test2"},
			expected:   []*gocrest.Matcher[string]{is.EqualTo("test1"), is.EqualTo("test2")},
			shouldFail: false,
		},
	}
	for _, test := range tests {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.EveryElement(test.expected...))

		then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
	}
}
func TestEveryIntElement(t *testing.T) {
	tests := []struct {
		actual     []int
		expected   []*gocrest.Matcher[int]
		shouldFail bool
	}{
		{
			actual:     []int{1, 2},
			expected:   []*gocrest.Matcher[int]{is.EqualTo(1), is.EqualTo(2)},
			shouldFail: false,
		},
		{
			actual: []int{1, 2},
			expected: []*gocrest.Matcher[int]{
				is.EqualTo(1),
			},
			shouldFail: true,
		},
		{
			actual: []int{1},
			expected: []*gocrest.Matcher[int]{
				is.EqualTo(1),
				is.EqualTo(2),
			},
			shouldFail: true,
		},
	}
	for _, test := range tests {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.EveryElement(test.expected...))

		then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
	}
}

func TestStructValues(t *testing.T) {
	tests := []struct {
		actual     any
		expected   has.StructMatchers[string]
		shouldFail bool
	}{
		{
			actual: struct {
				Id string
			}{Id: "Id"},
			expected: has.StructMatchers[string]{
				"Id": has.Prefix("Id"),
			},
			shouldFail: false,
		},
		{
			actual: struct {
				Id  string
				Id2 string
			}{Id: "Id", Id2: "Id2"},
			expected: has.StructMatchers[string]{
				"Id": has.Prefix("Id"),
			},
			shouldFail: false,
		},
		{
			actual: struct {
				Id string
			}{},
			expected: has.StructMatchers[string]{
				"Id": is.EmptyString(),
			},
			shouldFail: false,
		},
		{
			actual: struct {
				Id string
			}{},
			expected: has.StructMatchers[string]{
				"Id": is.EqualTo("something"),
			},
			shouldFail: true,
		},
		{
			actual: struct {
				Id  string
				Id2 string
			}{},
			expected: has.StructMatchers[string]{
				"Id2": is.EqualTo("something"),
			},
			shouldFail: true,
		},
		{
			actual: struct {
				Id  string
				Id2 string
			}{},
			expected: has.StructMatchers[string]{
				"Id":  is.EqualTo("Id"),
				"Id2": is.EqualTo("something"),
			},
			shouldFail: true,
		},
	}
	for _, test := range tests {
		stubTestingT := new(StubTestingT)
		then.AssertThat(stubTestingT, test.actual, has.StructWithValues[any](test.expected))

		then.AssertThat(t, stubTestingT.HasFailed(), is.EqualTo(test.shouldFail).Reason(stubTestingT.MockTestOutput))
	}
}

func TestConformsToStringer(t *testing.T) {
	then.AssertThat(t, is.Nil().String(), is.EqualTo("value that is <nil>"))
}

type DelayedReader struct {
	R io.Reader
	D time.Duration
}

func (s DelayedReader) Read(p []byte) (int, error) {
	time.Sleep(s.D)
	return s.R.Read(p)
}

func TestEventuallyWithDelayedReader(t *testing.T) {
	slowReader := DelayedReader{
		R: bytes.NewBuffer([]byte("abcdefghijklmnopqrstuv")),
		D: time.Second,
	}
	then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
		then.AssertThat(eventually, by.Reading(slowReader, 1024), is.EqualTo([]byte("abcdefghijklmnopqrstuv")))
	})
}

func TestEventuallyChannels(t *testing.T) {
	channel := firstTestChannel()
	then.Eventually(t, time.Second*5, time.Second, func(eventually gocrest.TestingT) {
		then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
	})
}

func TestEventuallyChannelsShouldFail(t *testing.T) {
	channel := firstTestChannel()
	channelTwo := secondTestChannel()
	stubbedTesting := new(StubTestingT)
	then.WithinTenSeconds(stubbedTesting, func(eventually gocrest.TestingT) {
		then.AssertThat(eventually, by.Channelling(channelTwo), is.EqualTo("11").Reason("This is going to fail"))
		then.AssertThat(eventually, by.Channelling(channel), is.EqualTo(3).Reason("should not fail"))
	})
	then.AssertThat(t, stubbedTesting.failed, is.EqualTo(true))
	then.AssertThat(t, stubbedTesting.MockTestOutput, is.AllOf(
		is.StringContaining("This is going to fail"),
		is.Not(is.StringContaining("should not fail")),
	))
}

func TestEventuallyChannelInterface(t *testing.T) {
	type MyType struct {
		F string
		B string
	}

	channel := make(chan *MyType, 1)
	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			m := new(MyType)
			m.F = fmt.Sprintf("hi - %d", i)
			m.B = fmt.Sprintf("bye - %d", i)
			channel <- m
		}
	}()
	then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
		then.AssertThat(eventually, by.Channelling(channel), has.StructWithValues[*MyType](has.StructMatchers[string]{
			"F": is.EqualTo("hi - 3"),
			"B": is.EqualTo("bye - 3"),
		}))
	})
}

func TestCallingFunctionEventually(t *testing.T) {
	function := func(a string) string {
		time.Sleep(time.Second)
		return a
	}
	then.WithinFiveSeconds(t, func(eventually gocrest.TestingT) {
		then.AssertThat(eventually, by.Calling(function, "hi"), is.EqualTo("hi"))
	})
}
func firstTestChannel() chan int {
	channel := make(chan int, 1)
	go func() {
		defer close(channel)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			channel <- i
		}
	}()
	return channel
}

func secondTestChannel() chan string {
	channelTwo := make(chan string, 1)
	go func() {
		defer close(channelTwo)
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			channelTwo <- fmt.Sprintf("%d", i)
		}
	}()
	return channelTwo
}
