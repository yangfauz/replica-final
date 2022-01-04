package migration

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	UserID               uint      `gorm:"not null"`
	RequestBy            string    `gorm:"not null"`
	Necessity            string    `gorm:"not null"`
	PaymentDate          time.Time `gorm:"not null"`
	PaymentAmount        int       `gorm:"not null"`
	PaymentCalculate     string    `gorm:"not null"`
	PaymentAccountName   string    `gorm:"not null"`
	PaymentAccountNumber string    `gorm:"not null"`
	StatusID             uint      `gorm:"not null"`
	Reason               string
}
