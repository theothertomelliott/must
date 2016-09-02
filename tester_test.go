package must

import (
	"errors"
	"reflect"
	"testing"
)

var beEqualTests = []struct {
	name       string
	expected   interface{}
	got        interface{}
	message    string
	shouldPass bool
	format     string
}{
	{
		name:       "Different strings",
		expected:   "string1",
		got:        "string2",
		shouldPass: false,
		format:     "%s: diff\n%s",
		message:    "Message1",
	},
	{
		name:       "Matching strings",
		expected:   "string",
		got:        "string",
		shouldPass: true,
	},
	{
		name:       "Different arrays",
		expected:   []string{"string1", "string2"},
		got:        []string{"string3", "string4"},
		shouldPass: false,
		format:     "%s: diff\n%s",
		message:    "Message2",
	},
	{
		name:       "Matching arrays",
		expected:   []string{"string1", "string2"},
		got:        []string{"string1", "string2"},
		shouldPass: true,
	},
}

func TestBeEqual(t *testing.T) {
	for _, test := range beEqualTests {
		m := &MockTesting{}
		tester := Tester{
			T: m,
		}
		result := tester.BeEqual(test.expected, test.got, test.message)
		if test.shouldPass && !result {
			t.Errorf("%s: Check did not pass as expected.", test.name)
		} else if !test.shouldPass && result {
			t.Errorf("%s: Check did not fail as expected", test.name)
		} else {
			if test.format != m.format {
				t.Errorf("%s: Incorrect error format. Expected '%v', got '%v'. errorCalled=%v", test.name, test.format, m.format, m.errorCalled)
			}

			if !result {
				if len(m.args) < 2 {
					t.Errorf("%s: Expected 2 error args, got %d", test.name, len(m.args))
				}

				if test.message != m.args[0] {
					t.Errorf("%s: Incorrect message. Expected '%v', got '%v'", test.name, test.message, m.args[0])
				}
			}
		}
	}
}

func TestBeEqualCustomCompare(t *testing.T) {
	for _, test := range beEqualTests {
		m := &MockTesting{}
		tester := Tester{
			T: m,
			InterfaceComparison: func(expected, got interface{}) bool {
				if !reflect.DeepEqual(expected, test.expected) {
					t.Errorf("Wrong expected sent to compare")
				}
				if !reflect.DeepEqual(got, test.got) {
					t.Errorf("Wrong got sent to compare")
				}
				return true
			},
		}
		if !tester.BeEqual(test.expected, test.got, test.message) {
			t.Errorf("Forced true comparison did not suceed as expected")
		}

		tester = Tester{
			T: m,
			InterfaceComparison: func(expected, got interface{}) bool {
				return false
			},
		}
		if tester.BeEqual(test.expected, test.got, test.message) {
			t.Errorf("Forced false comparison did not fail as expected")
		}
	}
}

func TestBeEqualCustomDiff(t *testing.T) {
	for _, test := range beEqualTests {
		m := &MockTesting{}
		tester := Tester{
			T: m,
			InterfaceDiff: func(expected, got interface{}) string {
				return "forced diff"
			},
		}
		result := tester.BeEqual(test.expected, test.got, test.message)
		if test.shouldPass && !result {
			t.Errorf("%s: Check did not pass as expected.", test.name)
		} else if !test.shouldPass && result {
			t.Errorf("%s: Check did not fail as expected", test.name)
		} else {
			if test.format != m.format {
				t.Errorf("%s: Incorrect error format. Expected '%v', got '%v'. errorCalled=%v", test.name, test.format, m.format, m.errorCalled)
			}

			if !result {
				if len(m.args) < 2 {
					t.Errorf("%s: Expected 2 error args, got %d", test.name, len(m.args))
				}

				if test.message != m.args[0] {
					t.Errorf("%s: Incorrect message. Expected '%v', got '%v'", test.name, test.message, m.args[0])
				}

				if "forced diff" != m.args[1] {
					t.Errorf("%s: Custom diff func was not used, got '%v'", test.name, m.args[1])
				}
			}
		}
	}
}

var beEqualErrorsTests = []struct {
	name       string
	expected   error
	got        error
	message    string
	shouldPass bool
	format     string
}{
	{
		name:       "Different errors",
		expected:   errors.New("Message one"),
		got:        errors.New("Message two"),
		shouldPass: false,
		format:     "%v\nExpected '%v', got '%v'",
	},
	{
		name:       "Matching errors",
		expected:   errors.New("Message"),
		got:        errors.New("Message"),
		shouldPass: true,
	},
}

func TestBeEqualErrors(t *testing.T) {
	for _, test := range beEqualErrorsTests {
		m := &MockTesting{}
		tester := Tester{
			T: m,
		}
		result := tester.BeEqualErrors(test.expected, test.got, test.message)
		if test.shouldPass && !result {
			t.Errorf("%s: Expected check would pass.", test.name)
		} else if !test.shouldPass && result {
			t.Errorf("%s: Expected check would not pass.", test.name)
		}

		if test.format != m.format {
			t.Errorf("%s: Incorrect error format. Expected '%v', got '%v'. errorCalled=%v", test.name, test.format, m.format, m.errorCalled)
		}
	}
}

var beNoErrorTest = []struct {
	name       string
	got        error
	message    string
	shouldPass bool
	format     string
}{
	{
		name:       "No error exists",
		shouldPass: true,
	},
	{
		name:       "Error exists",
		got:        errors.New("Message"),
		shouldPass: false,
		format:     "%s: error: %s",
	},
}

func TestBeNoError(t *testing.T) {
	for _, test := range beNoErrorTest {
		m := &MockTesting{}
		tester := Tester{
			T: m,
		}
		result := tester.BeNoError(test.got, test.message)
		if test.shouldPass && !result {
			t.Errorf("%s: Expected check would pass.", test.name)
		} else if !test.shouldPass && result {
			t.Errorf("%s: Expected check would not pass.", test.name)
		}

		if test.format != m.format {
			t.Errorf("%s: Incorrect error format. Expected '%v', got '%v'. errorCalled=%v", test.name, test.format, m.format, m.errorCalled)
		}
	}
}

var beSameLengthTests = []struct {
	name       string
	expected   interface{}
	got        interface{}
	message    string
	shouldPass bool
	format     string
}{
	{
		name:       "Strings, same length",
		expected:   "abcdefg",
		got:        "hijklmn",
		shouldPass: true,
	},
	{
		name:       "Strings, different length",
		expected:   "abc",
		got:        "defg",
		shouldPass: false,
		format:     "%s: expected length %d, got length %d",
	},
	{
		name:       "Arrays, same length",
		expected:   []int{1, 2, 3},
		got:        []int{4, 5, 6},
		shouldPass: true,
	},
	{
		name:       "Arrays, different length",
		expected:   []int{1, 2, 3, 7},
		got:        []int{8},
		shouldPass: false,
		format:     "%s: expected length %d, got length %d",
	},
	{
		name:       "String and string pointer, same length",
		expected:   stringToPointer("abcdefg"),
		got:        "hijklmn",
		shouldPass: true,
	},
	{
		name:       "String and string pointer, different length",
		expected:   "abc",
		got:        stringToPointer("defg"),
		shouldPass: false,
		format:     "%s: expected length %d, got length %d",
	},
	{
		name:       "String and struct",
		expected:   "abc",
		got:        struct{ content string }{content: "test"},
		shouldPass: false,
		format:     "%s: could not test lengths - %v",
	},
}

func TestBeSameLength(t *testing.T) {
	for _, test := range beSameLengthTests {
		m := &MockTesting{}
		tester := Tester{
			T: m,
		}
		result := tester.BeSameLength(test.expected, test.got, test.message)
		if test.shouldPass && !result {
			t.Errorf("%s: Check did not pass as expected.", test.name)
		} else if !test.shouldPass && result {
			t.Errorf("%s: Check did not fail as expected", test.name)
		}

		if test.format != m.format {
			t.Errorf("%s: Incorrect error format. Expected '%v', got '%v'. args=%v, errorCalled=%v", test.name, test.format, m.format, m.args, m.errorCalled)
		}

		if !result {
			if len(m.args) < 2 {
				t.Errorf("%s: Expected at least 2 error args, got %d", test.name, len(m.args))
			}

			if test.message != m.args[0] {
				t.Errorf("%s: Incorrect message. Expected '%v', got '%v'", test.name, test.message, m.args)
			}
		}
	}
}

func stringToPointer(val string) *string {
	return &val
}

type MockTesting struct {
	errorCalled bool
	format      string
	args        []interface{}
}

func (m *MockTesting) Errorf(format string, args ...interface{}) {
	m.errorCalled = true
	m.format = format
	m.args = args
}
