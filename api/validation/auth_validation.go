package validation

import (
	"replica-finalproject/api/model"
	"replica-finalproject/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func LoginValidate(request model.UserLoginRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required.When(request.Username == "").Error("Username is Required")),
		validation.Field(&request.Password, validation.Required.When(request.Password == "").Error("Password is Required")),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
