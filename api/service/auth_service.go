package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
)

type AuthService interface {
	Login(input model.UserLoginRequest) (entity.User, error)
}
