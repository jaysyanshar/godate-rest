package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	AccountID uint      `json:"account_id" gorm:"not null;index"`
	Account   Account   `gorm:"foreignKey:AccountID;references:ID"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
}
