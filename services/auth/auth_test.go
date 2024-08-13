package auth

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jaysyanshar/godate-rest/config"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
	"github.com/jaysyanshar/godate-rest/models/restmodel"
	"github.com/jaysyanshar/godate-rest/repositories/account"
	"github.com/jaysyanshar/godate-rest/repositories/profile"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	cfg         *config.Config
	accountRepo *account.MockAccountRepository
	profileRepo *profile.MockProfileRepository
	authService AuthService
)

func setup(t *testing.T) {
	ctrl := gomock.NewController(t)
	accountRepo = account.NewMockAccountRepository(ctrl)
	profileRepo = profile.NewMockProfileRepository(ctrl)

	cfg = &config.Config{
		AppName:   "GoDate",
		JwtSecret: "jwtsecret",
	}

	authService = NewService(cfg, accountRepo, profileRepo)
}

func teardown() {
	cfg = nil
	accountRepo = nil
	profileRepo = nil
	authService = nil
}

func TestService_SignUp(t *testing.T) {
	setup(t)
	defer teardown()

	ctx := context.Background()

	tests := []struct {
		name          string
		req           restmodel.SignUpRequest
		mockSetup     func()
		res           restmodel.SignUpResponse
		expectedError bool
	}{
		{
			name: "successful signup",
			req: restmodel.SignUpRequest{
				Email:     "testuser@mail.com",
				Password:  "testpassword",
				FirstName: "John",
				LastName:  "Doe",
				BirthDate: "1990-01-01",
				Gender:    "male",
			},
			res: restmodel.SignUpResponse{
				Success: true,
				Message: "Account created successfully",
			},
			mockSetup: func() {
				accountRepo.EXPECT().Insert(ctx, gomock.Any()).Return(uint(1), nil)
				profileRepo.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(uint(1), nil)
			},
			expectedError: false,
		},
		{
			name: "invalid request",
			req:  restmodel.SignUpRequest{},
			res: restmodel.SignUpResponse{
				Success: false,
			},
			mockSetup:     func() {},
			expectedError: true,
		},
		{
			name: "account insert failed",
			req: restmodel.SignUpRequest{
				Email:     "testuser@mail.com",
				Password:  "testpassword",
				FirstName: "John",
				LastName:  "Doe",
				BirthDate: "1990-01-01",
				Gender:    "male",
			},
			res: restmodel.SignUpResponse{
				Success: false,
			},
			mockSetup: func() {
				accountRepo.EXPECT().Insert(ctx, gomock.Any()).Return(uint(0), assert.AnError)
			},
			expectedError: true,
		},
		{
			name: "profile insert failed",
			req: restmodel.SignUpRequest{
				Email:     "testuser@mail.com",
				Password:  "testpassword",
				FirstName: "John",
				LastName:  "Doe",
				BirthDate: "1990-01-01",
				Gender:    "male",
			},
			res: restmodel.SignUpResponse{
				Success: false,
			},
			mockSetup: func() {
				accountRepo.EXPECT().Insert(ctx, gomock.Any()).Return(uint(1), nil)
				profileRepo.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(uint(0), assert.AnError)
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := authService.SignUp(ctx, tt.req)
			if (err != nil) != tt.expectedError {
				t.Errorf("SignUp() error = %v, expectedError %v", err, tt.expectedError)
			}
			assert.Equal(t, tt.res.Success, res.Success)
		})
	}
}

func TestService_Login(t *testing.T) {
	setup(t)
	defer teardown()

	ctx := context.Background()
	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	tests := []struct {
		name          string
		req           restmodel.LoginRequest
		mockSetup     func()
		res           restmodel.LoginResponse
		expectedError bool
	}{
		{
			name: "successful login",
			req: restmodel.LoginRequest{
				Email:    "test@mail.com",
				Password: "password",
			},
			res: restmodel.LoginResponse{
				Success: true,
				Message: "Login successful",
				Token:   "jwtToken",
			},
			mockSetup: func() {
				accountRepo.EXPECT().FindByEmail(ctx, "test@mail.com").Return(dbmodel.Account{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@mail.com",
					Password: string(encryptedPass),
				}, nil)

			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := authService.Login(ctx, tt.req)
			if (err != nil) != tt.expectedError {
				t.Errorf("Login() error = %v, expectedError %v", err, tt.expectedError)
			}
			assert.Equal(t, tt.res.Success, res.Success)
			if !tt.expectedError {
				assert.True(t, res.Token != "")
			}
		})
	}
}
