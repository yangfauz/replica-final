package repository

import "replica-finalproject/api/entity"

type RoleRepository interface {
	FindAll() ([]entity.Role, error)
	FindById(id int) (entity.Role, error)
	FindByName(name string) (entity.Role, error)
	Save(role entity.Role) (entity.Role, error)
	Update(id int, role entity.Role) (entity.Role, error)
	Delete(id int) (entity.Role, error)
}
