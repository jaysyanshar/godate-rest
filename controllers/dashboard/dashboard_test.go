package dashboard

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a new instance of the controller
	ctrl := NewController()

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP response recorder
	res := httptest.NewRecorder()

	// Call the HelloHandler method
	ctrl.HelloHandler(res, req)

	// Check the response status code
	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}

	// Check the response body
	expectedBody := "Hello, World!"
	if res.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, res.Body.String())
	}
}
