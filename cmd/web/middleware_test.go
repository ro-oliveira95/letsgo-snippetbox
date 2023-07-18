package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ro-oliveira95/letsgo-snippetbox/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	// Initialize a new httptest.ResponseRecorder and dummy http.Request.
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP handler to test if it is properly called.
	mNext := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Call middleware with mocked next handler and call returned
	// handler's ServeHTTP method with our ResponseRecorder and dummy Request.
	secureHeaders(mNext).ServeHTTP(rec, req)
	res := rec.Result()

	// Test headers individually.
	headers := []struct {
		k string // Header key
		v string // Expected value
	}{
		{k: "Content-Security-Policy", v: "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"},
		{k: "Referrer-Policy", v: "origin-when-cross-origin"},
		{k: "X-Content-Type-Options", v: "nosniff"},
		{k: "X-Frame-Options", v: "deny"},
		{k: "X-XSS-Protection", v: "0"},
	}

	for _, h := range headers {
		assert.Equal(t, res.Header.Get(h.k), h.v, h.k+" header")
	}

	// Check response's status code and body content.
	assert.Equal(t, res.StatusCode, http.StatusOK, "")
	assert.BodyContains(t, res, "OK")
}
