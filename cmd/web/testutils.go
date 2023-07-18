package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// Helper function newTestApplication returns an instance of our
// application struct containing mocked dependencies.
func newTestApplication(t *testing.T) *application {
	return &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}
}

// Custom type which embeds a httptest.Server instance.
type testServer struct {
	*httptest.Server
}

// Helper function newTestServer initalizes and returns a new instance
// of our custom testServer type.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	// Initialize a new cookie jar and add it to our server client
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	ts.Client().Jar = jar

	// Prevent client from redirections
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// Method get makes a GET request to a given url path using the test
// server client, and returns the response status code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	res, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}
	body := readResponseBody(t, res)
	return res.StatusCode, res.Header, body
}

// Function readResponseBody reads the body of a HTTP response
// and return it as string. If an error occurs it will call the
// testing runtime to abort the execution.
func readResponseBody(t *testing.T, r *http.Response) string {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)
	return string(body)
}
