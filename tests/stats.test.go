package tests

// TODO: all

import (
	"net/http"
	"net/http/httptest"
	"r314dash/handler"
	"testing"
)

func TestStatsHandler(t *testing.T) {
	// Create a new request to pass to our handler.
	req := httptest.NewRequest("GET", "http://example.com/stats", nil)

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Call the handler, passing in the request and the ResponseRecorder.
	handler := http.HandlerFunc(handler.StatsHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "<h1>Hello, World!</h1>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// import (
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// // Mock HTTP client for testing
// type mockTransport struct {
// 	response *http.Response
// 	err      error
// }

// func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	return m.response, m.err
// }

// func TestMyHandler(t *testing.T) {
// 	// Create a new request to pass to our handler.
// 	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

// 	// Create a ResponseRecorder to record the response.
// 	rr := httptest.NewRecorder()

// 	// Set up the mock response
// 	mockResponse := &http.Response{
// 		StatusCode: http.StatusOK,
// 		Body:       ioutil.NopCloser(strings.NewReader("mock response")),
// 	}

// 	// Replace the global client with a mock client
// 	client = &http.Client{
// 		Transport: &mockTransport{
// 			response: mockResponse,
// 		},
// 	}

// 	// Call the handler, passing in the request and the ResponseRecorder.
// 	handler := http.HandlerFunc(MyHandler)
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body is what we expect.
// 	expected := "<h1>Hello, World!</h1>"
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
