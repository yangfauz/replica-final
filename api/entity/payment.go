package entity

import "time"

type Payment struct {
	ID                   uint
	UserID               uint
	RequestBy            string
	Necessity            string
	PaymentDate          time.Time
	PaymentAmount        int
	PaymentCalculate     string
	PaymentAccountName   string
	PaymentAccountNumber string
	StatusID             uint
	Reason               string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
