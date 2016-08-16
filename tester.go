package must

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
	return BeEqual(tester.T, expected, got, message)
}

/*
BeEqualErrors compares the expected and got errors, triggering an error on the Tester's T if they are not equal.

This corresponds to the function BeEqualErrors
*/
func (tester Tester) BeEqualErrors(expected, got error, message string) bool {
	return BeEqualErrors(tester.T, expected, got, message)
}
