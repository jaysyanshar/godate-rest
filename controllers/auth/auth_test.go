package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaysyanshar/godate-rest/models/restmodel"
	"github.com/jaysyanshar/godate-rest/services/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := auth.NewMockAuthService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name           string
		req            restmodel.SignUpRequest
		mockSetup      func()
		expectedResult restmodel.SignUpResponse
		expectedError  bool
	}{
		{
			name: "successful signup",
			req: restmodel.SignUpRequest{
				Email:    "testuser@mail.com",
				Password: "testpassword",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().SignUp(ctx, gomock.Any()).Return(restmodel.SignUpResponse{
					Success: true,
				}, nil)
			},
			expectedResult: restmodel.SignUpResponse{
				Success: true,
			},
			expectedError: false,
		},
		{
			name: "signup failure",
			req: restmodel.SignUpRequest{
				Email:    "testuser2@mail.com",
				Password: "testpassword2",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().SignUp(ctx, gomock.Any()).Return(restmodel.SignUpResponse{}, errors.New("signup failed"))
			},
			expectedResult: restmodel.SignUpResponse{
				Success: false,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := mockAuthService.SignUp(ctx, tt.req)
			if (err != nil) != tt.expectedError {
				t.Errorf("SignUp() error = %v, expectedError %v", err, tt.expectedError)
			}
			assert.Equal(t, tt.expectedResult.Success, res.Success)
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := auth.NewMockAuthService(ctrl)

	ctx := context.Background()

	tests := []struct {
		name           string
		req            restmodel.LoginRequest
		mockSetup      func()
		expectedResult restmodel.LoginResponse
		expectedError  bool
	}{
		{
			name: "successful login",
			req: restmodel.LoginRequest{
				Email:    "testuser@mail.com",
				Password: "testpassword",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().Login(ctx, gomock.Any()).Return(restmodel.LoginResponse{
					Success: true,
					Token:   "token",
				}, nil)
			},
			expectedResult: restmodel.LoginResponse{
				Success: true,
				Token:   "token",
			},
			expectedError: false,
		},
		{
			name: "login failure",
			req: restmodel.LoginRequest{
				Email:    "testuser2@mail.com",
				Password: "testpassword2",
			},
			mockSetup: func() {
				mockAuthService.EXPECT().Login(ctx, gomock.Any()).Return(restmodel.LoginResponse{}, errors.New("login failed"))
			},
			expectedResult: restmodel.LoginResponse{
				Success: false,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := mockAuthService.Login(ctx, tt.req)
			if (err != nil) != tt.expectedError {
				t.Errorf("Login() error = %v, expectedError %v", err, tt.expectedError)
			}
			assert.Equal(t, tt.expectedResult.Success, res.Success)
			if !tt.expectedError {
				assert.Equal(t, tt.expectedResult.Token, res.Token)
			}
		})
	}
}
