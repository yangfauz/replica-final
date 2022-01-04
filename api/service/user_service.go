package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
)

type UserService interface {
	GetAll() ([]entity.User, error)
	GetById(id int) (entity.User, error)
	Create(input model.CreateUserRequest) (entity.User, error)
	Update(id int, input model.CreateUserRequest) (entity.User, error)
	Delete(id int) bool
}
