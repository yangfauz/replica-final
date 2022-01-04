package validation

import (
	"replica-finalproject/api/model"
	"replica-finalproject/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func PaymentValidate(request model.CreatePaymentRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.RequestBy, validation.Required.When(request.RequestBy == "").Error("Request By is Required")),
		validation.Field(&request.Necessity, validation.Required.When(request.Necessity == "").Error("Necessity is Required")),
		validation.Field(&request.PaymentAmount, validation.Required, validation.Min(0)),
		validation.Field(&request.PaymentAccountName, validation.Required.When(request.PaymentAccountName == "").Error("Payment Account Name is Required")),
		validation.Field(&request.PaymentAccountNumber, validation.Required.When(request.PaymentAccountNumber == "").Error("Payment Account Number is Required")),
		validation.Field(&request.StatusID, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
