package must

/*
BeEqual compares the expected and got interfaces, triggering an error on t if they are not equal.
This error will include a diff of the two objects.

The return value will be true if the interfaces are equal.

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeEqual(t TestingT, expected, got interface{}, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeEqual(expected, got, a...)
}

/*
BeEqualErrors compares two errors to determine if they are considered equal.
The errors expected and got are considered equal if they are both nil, or both are non-nil and their error messsages (from their Error() functions) match.

This ignores the actual type of these errors, so two errors created with different struct types, but the same message will still be equal.

Should the errors not be considered equal, an error will be raised in t including both messages and false will be returned.

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeEqualErrors(t TestingT, expected, got error, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeEqualErrors(expected, got, a...)
}

/*
BeNoError checks whether or not the got value is an error.

The return value will be true if got is nil.

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeNoError(t TestingT, got error, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeNoError(got, a...)
}

/*
BeSameLength checks whether the two inputs have the same length according to the len function.

The return value will be true if their lengths match.

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeSameLength(t TestingT, expected, got interface{}, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeSameLength(expected, got, a...)
}

/*
BeError checks that the provided error is not nil

The return value will be true iff error is not nil

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeError(t TestingT, got error, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeError(got, a...)
}

/*
BeErrorIf checks that the errorExpected flag corresponds to the provided error: error should be
nil iff errorExpected flag is false

The return value will be true iff error value corresponds to the errorExpected flag

Additional output for any error message can be provided as additional parameters, as with fmt.Print.
*/
func BeErrorIf(t TestingT, errorExpected bool, got error, a ...interface{}) bool {
	t.Helper()
	mt := Tester{T: t}
	return mt.BeErrorIf(errorExpected, got, a...)
}
