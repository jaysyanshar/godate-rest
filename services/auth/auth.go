package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jaysyanshar/godate-rest/config"
	"github.com/jaysyanshar/godate-rest/models/restmodel"
	"github.com/jaysyanshar/godate-rest/repositories/account"
	"github.com/jaysyanshar/godate-rest/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(ctx context.Context, req restmodel.SignUpRequest) (restmodel.SignUpResponse, error)
	Login(ctx context.Context, req restmodel.LoginRequest) (restmodel.LoginResponse, error)
}

type service struct {
	cfg         *config.Config
	accountRepo account.AccountRepository
	userRepo    user.UserRepository
}

func NewService(cfg *config.Config, accountRepo account.AccountRepository, userRepo user.UserRepository) AuthService {
	return &service{
		cfg:         cfg,
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}

func (s *service) SignUp(ctx context.Context, req restmodel.SignUpRequest) (restmodel.SignUpResponse, error) {
	var err error
	if err = req.Validate(); err != nil {
		return restmodel.SignUpResponse{Success: false, Message: err.Error()}, err
	}

	// encrypt password
	account := req.ToAccount()
	password, err := encryptPassword(account.Password)
	if err != nil {
		return restmodel.SignUpResponse{Success: false, Message: err.Error()}, err
	}
	account.Password = password

	// Insert account
	accountID, err := s.accountRepo.Insert(ctx, account)
	if err != nil {
		return restmodel.SignUpResponse{Success: false, Message: err.Error()}, err
	}

	// Insert user
	user := req.ToUser(accountID)
	_, err = s.userRepo.Insert(ctx, user)
	if err != nil {
		return restmodel.SignUpResponse{Success: false, Message: err.Error()}, err
	}

	return restmodel.SignUpResponse{Success: true, Message: "Account created successfully"}, nil
}

func (s *service) Login(ctx context.Context, req restmodel.LoginRequest) (restmodel.LoginResponse, error) {
	account, err := s.accountRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return restmodel.LoginResponse{Success: false, Message: err.Error()}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return restmodel.LoginResponse{Success: false, Message: fmt.Sprintf("Invalid email or password: %v", err)}, err
	}

	jwtToken, err := generateJWTToken(account.ID, s.cfg.AppName, s.cfg.JwtSecret)
	if err != nil {
		return restmodel.LoginResponse{Success: false, Message: fmt.Sprintf("Failed to generate JWT Token: %v", err)}, err
	}
	return restmodel.LoginResponse{Success: true, Message: "Login successful", Token: jwtToken}, nil
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt password: %w", err)
	}
	encryptedPassword := string(hashedPassword)
	return encryptedPassword, nil
}

func generateJWTToken(accountID uint, issuer, jwtKey string) (string, error) {
	// Define token claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    issuer,
		Subject:   fmt.Sprint(accountID),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Convert the jwtKey to a byte slice
	key := []byte(jwtKey)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
