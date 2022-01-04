package validation

import (
	"replica-finalproject/api/model"
	"replica-finalproject/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func UserValidate(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
		validation.Field(&request.Username, validation.Required.When(request.Username == "").Error("Username is Required")),
		validation.Field(&request.Password, validation.Required.When(request.Password == "").Error("Password is Required")),
		validation.Field(&request.RoleID, validation.Required, validation.Min(0)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
