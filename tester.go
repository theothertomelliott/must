package must

import "github.com/kylelemons/godebug/pretty"

var _ MustTester = Tester{}

/*
Tester implements MustTester and provides a TestingT to be used for all check functions.
*/
type Tester struct {
	T                   TestingT                               // *testing.T or equivalent
	InterfaceComparison func(expected, got interface{}) bool   // Optional custom interface comparison function
	InterfaceDiff       func(expected, got interface{}) string // Optional custom interace diff function
}

/*
BeEqual compares the expected and got interfaces, triggering an error on the Tester's T if they are not equal.

This corresponds to the function BeEqual
*/
func (tester Tester) BeEqual(expected, got interface{}, message string) bool {
	if !tester.equal(expected, got) {
		tester.T.Errorf("%s: diff: (-got +want)\n%s", message, tester.diff(expected, got))
		return false
	}
	return true
}

/*
BeEqualErrors compares the expected and got errors, triggering an error on the Tester's T if they are not equal.

This corresponds to the function BeEqualErrors
*/
func (tester Tester) BeEqualErrors(expected, got error, message string) bool {
	if expected == nil && got == nil {
		return true
	}
	if (expected == nil || got == nil) || expected.Error() != got.Error() {
		tester.T.Errorf("%v\nExpected '%v', got '%v'", message, getErrMessage(expected), getErrMessage(got))
		return false
	}
	return true
}

/*
BeNoError checks whether got is set, triggering an error on the Tester's T if it is non-nil.

This corresponds to the function BeNoError
*/
func (tester Tester) BeNoError(got error, message string) bool {
	if got == nil {
		return true
	}
	tester.T.Errorf("%s: error: %s", message, got.Error())
	return false
}

func (tester Tester) equal(expected, got interface{}) bool {
	if tester.InterfaceComparison != nil {
		return tester.InterfaceComparison(expected, got)
	}
	return pretty.Compare(expected, got) == ""
}

func (tester Tester) diff(expected, got interface{}) string {
	if tester.InterfaceDiff != nil {
		return tester.InterfaceDiff(expected, got)
	}
	return pretty.Compare(expected, got)
}

func getErrMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return "<nil>"
}
