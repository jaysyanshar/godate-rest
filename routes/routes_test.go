// FILEPATH: /Users/jaysyansharulloh/Documents/code-practice/godate-rest/routes/routes_test.go

package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaysyanshar/godate-rest/controllers/auth"
	"github.com/jaysyanshar/godate-rest/controllers/dashboard"
	"github.com/jaysyanshar/godate-rest/middlewares"
)

func TestSetupRouter(t *testing.T) {
	// Create mock instances of the dependencies
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMiddleware := middlewares.NewMockMiddleware(ctrl)
	mockAuthController := auth.NewMockAuthController(ctrl)
	mockDashboardController := dashboard.NewMockDashboardController(ctrl)

	// Setup expectations for the mocks
	mockMiddleware.EXPECT().JWTMiddleware(gomock.Any()).Return(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})).AnyTimes()

	mockAuthController.EXPECT().LoginHandler(gomock.Any(), gomock.Any()).Return().AnyTimes()
	mockDashboardController.EXPECT().HelloHandler(gomock.Any(), gomock.Any()).Return().AnyTimes()

	// Initialize the router with the mocked dependencies
	router := SetupRouter(mockMiddleware, mockAuthController, mockDashboardController)

	// Define test cases
	tests := []struct {
		method       string
		url          string
		expectedCode int
	}{
		{"GET", "/", http.StatusOK},
		{"POST", "/api/v1/login", http.StatusOK},
	}

	// Run test cases
	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.url, nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		if res.Code != tt.expectedCode {
			t.Errorf("expected status code %d, got %d", tt.expectedCode, res.Code)
		}
	}
}
