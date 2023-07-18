package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ro-oliveira95/letsgo-snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// Initialize a new httptest.ResponseRecorder and dummy http.Request.
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the ping handler function and retrieve result
	ping(rec, req)
	res := rec.Result()

	// Check response's status code and body content
	assert.Equal(t, res.StatusCode, http.StatusOK, "")
	assert.BodyContains(t, res, "OK")
}
