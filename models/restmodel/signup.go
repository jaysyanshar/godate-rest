package restmodel

import (
	"fmt"
	"strings"
	"time"

	"github.com/jaysyanshar/godate-rest/models/constant"
	"github.com/jaysyanshar/godate-rest/models/dbmodel"
)

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Gender    string `json:"gender"`
}

type SignUpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (req SignUpRequest) ToAccount() dbmodel.Account {
	return dbmodel.Account{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req SignUpRequest) ToUser(accountID uint) dbmodel.User {
	birthDate, _ := time.Parse("2006-01-02", req.BirthDate)
	return dbmodel.User{
		AccountID: accountID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: birthDate,
		Gender:    req.Gender,
	}
}

func (req SignUpRequest) Validate() error {
	if req.Email == "" {
		return fmt.Errorf(constant.ErrEmptyEmail)
	}
	if req.Password == "" {
		return fmt.Errorf(constant.ErrEmptyPassword)
	}
	if req.FirstName == "" {
		return fmt.Errorf(constant.ErrEmptyFirstName)
	}
	if req.LastName == "" {
		return fmt.Errorf(constant.ErrEmptyLastName)
	}
	if !isValidDateFormat(req.BirthDate) {
		return fmt.Errorf(constant.ErrInvalidBirthDate)
	}
	if !strings.EqualFold(req.Gender, "male") && !strings.EqualFold(req.Gender, "female") {
		return fmt.Errorf(constant.ErrInvalidGender)
	}
	return nil
}

func isValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
