package dbmodel

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
