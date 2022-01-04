package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
)

type RoleService interface {
	GetAll() ([]entity.Role, error)
	GetById(id int) (entity.Role, error)
	Create(input model.CreateRoleRequest) (entity.Role, error)
	Update(id int, input model.CreateRoleRequest) (entity.Role, error)
	Delete(id int) bool
}
