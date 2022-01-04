package model

import (
	"replica-finalproject/api/entity"
	"time"

	"github.com/leekchan/accounting"
)

//request
type CreatePaymentRequest struct {
	RequestBy            string    `json:"request_by" binding:"required"`
	Necessity            string    `json:"necessity" binding:"required"`
	PaymentDate          string    `json:"payment_date" binding:"required"`
	PaymentAmount        int       `json:"payment_amount" binding:"required"`
	PaymentCalculate     string    `json:"payment_calculate" binding:"required"`
	PaymentAccountName   string    `json:"payment_account_name" binding:"required"`
	PaymentAccountNumber string    `json:"payment_account_number" binding:"required"`
	StatusID             uint      `json:"status_id" binding:"required"`
	Reason               string    `json:"reason" binding:"required"`
	CreatedAt            time.Time `json:"created_at" binding:"required"`
	UpdatedAt            time.Time `json:"updated_at" binding:"required"`
}

//response
type GetPaymentResponse struct {
	ID                   uint      `json:"id" binding:"required"`
	UserID               uint      `json:"unit_id" binding:"required"`
	RequestBy            string    `json:"request_by" binding:"required"`
	Necessity            string    `json:"necessity" binding:"required"`
	PaymentDate          time.Time `json:"payment_date" binding:"required"`
	PaymentAmount        int       `json:"payment_amount" binding:"required"`
	PaymentCalculate     string    `json:"payment_calculate" binding:"required"`
	PaymentAccountName   string    `json:"payment_account_name" binding:"required"`
	PaymentAccountNumber string    `json:"payment_account_number" binding:"required"`
	StatusID             uint      `json:"status_id" binding:"required"`
	Reason               string    `json:"reason" binding:"required"`
	CreatedAt            time.Time `json:"created_at" binding:"required"`
	UpdatedAt            time.Time `json:"updated_at" binding:"required"`
}

type CreatePaymentResponse struct {
	ID                   uint      `json:"id" binding:"required"`
	UserID               uint      `json:"unit_id" binding:"required"`
	RequestBy            string    `json:"request_by" binding:"required"`
	Necessity            string    `json:"necessity" binding:"required"`
	PaymentDate          time.Time `json:"payment_date" binding:"required"`
	PaymentAmount        int       `json:"payment_amount" binding:"required"`
	PaymentCalculate     string    `json:"payment_calculate" binding:"required"`
	PaymentAccountName   string    `json:"payment_account_name" binding:"required"`
	PaymentAccountNumber string    `json:"payment_account_number" binding:"required"`
	StatusID             uint      `json:"status_id" binding:"required"`
	Reason               string    `json:"reason" binding:"required"`
	CreatedAt            time.Time `json:"created_at" binding:"required"`
	UpdatedAt            time.Time `json:"updated_at" binding:"required"`
}

func FormatGetAllPaymentResponse(payments []entity.Payment) []GetPaymentResponse {
	paymentsFormatter := []GetPaymentResponse{}

	for _, payment := range payments {
		paymentFormatter := GetPaymentResponse(payment)
		paymentsFormatter = append(paymentsFormatter, paymentFormatter)
	}

	return paymentsFormatter
}

func FormatGetPaymentResponse(payment entity.Payment) GetPaymentResponse {

	paymentFormatter := GetPaymentResponse{}
	paymentFormatter.ID = payment.ID
	paymentFormatter.UserID = payment.UserID
	paymentFormatter.RequestBy = payment.RequestBy
	paymentFormatter.Necessity = payment.Necessity
	paymentFormatter.PaymentDate = payment.PaymentDate
	paymentFormatter.PaymentAmount = payment.PaymentAmount
	paymentFormatter.PaymentCalculate = payment.PaymentCalculate
	paymentFormatter.PaymentAccountName = payment.PaymentAccountName
	paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber
	paymentFormatter.StatusID = payment.StatusID
	paymentFormatter.Reason = payment.Reason
	paymentFormatter.CreatedAt = payment.CreatedAt
	paymentFormatter.UpdatedAt = payment.UpdatedAt

	return paymentFormatter
}

func FormatCreatePaymentResponse(payment entity.Payment) CreatePaymentResponse {

	paymentFormatter := CreatePaymentResponse{}
	paymentFormatter.ID = payment.ID
	paymentFormatter.UserID = payment.UserID
	paymentFormatter.RequestBy = payment.RequestBy
	paymentFormatter.Necessity = payment.Necessity
	paymentFormatter.PaymentDate = payment.PaymentDate
	paymentFormatter.PaymentAmount = payment.PaymentAmount
	paymentFormatter.PaymentCalculate = payment.PaymentCalculate
	paymentFormatter.PaymentAccountName = payment.PaymentAccountName
	paymentFormatter.PaymentAccountNumber = payment.PaymentAccountNumber
	paymentFormatter.StatusID = payment.StatusID
	paymentFormatter.Reason = payment.Reason
	paymentFormatter.CreatedAt = payment.CreatedAt
	paymentFormatter.UpdatedAt = payment.UpdatedAt

	return paymentFormatter
}

func (payment GetPaymentResponse) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(payment.PaymentAmount)
}
