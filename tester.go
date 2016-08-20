package must

import "github.com/kylelemons/godebug/pretty"

var _ MustTester = Tester{}

/*
Tester implements MustTester and provides a TestingT to be used for all check functions.
*/
type Tester struct {
	T TestingT
}

/*
BeEqual compares the expected and got interfaces, triggering an error on the Tester's T if they are not equal.

This corresponds to the function BeEqual
*/
func (tester Tester) BeEqual(expected, got interface{}, message string) bool {
	if diff := pretty.Compare(expected, got); diff != "" {
		tester.T.Errorf("%s: diff: (-got +want)\n%s", message, diff)
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

func getErrMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return "<nil>"
}
