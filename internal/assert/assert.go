package assert

import (
	"bytes"
	"io"
	"net/http"
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

// BodyContains testing helper function reads buffered body
// from http.Response and asserts its stringified version
// matches the value argument.
//
// Attention: after calling this function the Response's body
// will be empty, so to use this function again another response
// body buffer must be fullfilled again.
func BodyContains(t *testing.T, res *http.Response, value string) {
	t.Helper()

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	Equal(t, string(body), value, "")
}
