package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T, prefix string) {
	t.Helper()

	if actual != expected {
		message := "got: %v; want: %v"
		if prefix != "" {
			message = prefix + ": " + message
		}
		t.Errorf(message, actual, expected)
	}
}
