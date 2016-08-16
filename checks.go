package must

import "github.com/kylelemons/godebug/pretty"

/*
BeEqual compares the expected and got interfaces, triggering an error on t if they are not equal.
This error will include a diff of the two objects.

The return value will be true if the interfaces are equal.
*/
func BeEqual(t TestingT, expected, got interface{}, message string) bool {
	if diff := pretty.Compare(expected, got); diff != "" {
		t.Errorf("%s: diff: (-got +want)\n%s", message, diff)
		return false
	}
	return true
}

/*
BeEqualErrors compares two errors to determine if they are considered equal.
The errors expected and got are considered equal if they are both nil, or both are non-nil and their error messsages (from their Error() functions) match.

This ignores the actual type of these errors, so two errors created with different struct types, but the same message will still be equal.

Should the errors not be considered equal, an error will be raised in t including both messages and false will be returned.
*/
func BeEqualErrors(t TestingT, expected, got error, message string) bool {
	if expected == nil && got == nil {
		return true
	}
	if (expected == nil || got == nil) || expected.Error() != got.Error() {
		t.Errorf("%v\nExpected '%v', got '%v'", message, getErrMessage(expected), getErrMessage(got))
		return false
	}
	return true
}

func getErrMessage(err error) string {
	if err != nil {
		return err.Error()
	}
	return "<nil>"
}
