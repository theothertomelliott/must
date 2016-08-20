package must

/*
BeEqual compares the expected and got interfaces, triggering an error on t if they are not equal.
This error will include a diff of the two objects.

The return value will be true if the interfaces are equal.
*/
func BeEqual(t TestingT, expected, got interface{}, message string) bool {
	mt := Tester{T: t}
	return mt.BeEqual(expected, got, message)
}

/*
BeEqualErrors compares two errors to determine if they are considered equal.
The errors expected and got are considered equal if they are both nil, or both are non-nil and their error messsages (from their Error() functions) match.

This ignores the actual type of these errors, so two errors created with different struct types, but the same message will still be equal.

Should the errors not be considered equal, an error will be raised in t including both messages and false will be returned.
*/
func BeEqualErrors(t TestingT, expected, got error, message string) bool {
	mt := Tester{T: t}
	return mt.BeEqualErrors(expected, got, message)
}

/*
BeNoError checks whether or not the got value is an error.

The return value will be true if got is nil.
*/
func BeNoError(t TestingT, got error, message string) bool {
	mt := Tester{T: t}
	return mt.BeNoError(got, message)
}
