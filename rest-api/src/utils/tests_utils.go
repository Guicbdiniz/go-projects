package utils

import "testing"

func CheckTestError(t *testing.T, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %s", message, err.Error())
	}
}

func AssertEqual[V comparable](t *testing.T, expected V, actual V, message string) {
	t.Helper()

	if expected != actual {
		t.Fatalf("%s. Expected %v, got %v", message, expected, actual)
	}
}
