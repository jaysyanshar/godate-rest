package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountID uint      `json:"account_id" gorm:"not null;index"`
	Account   Account   `gorm:"foreignKey:AccountID;references:ID"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	Gender    string    `json:"gender" gorm:"not null"`
}
