package dbmodel

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
