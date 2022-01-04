package validation

import (
	"replica-finalproject/api/model"
	"replica-finalproject/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func RoleValidate(request model.CreateRoleRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
