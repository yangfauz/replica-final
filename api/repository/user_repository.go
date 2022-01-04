package repository

import "replica-finalproject/api/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	Save(user entity.User) (entity.User, error)
	Update(id int, user entity.User) (entity.User, error)
	Delete(id int) (entity.User, error)
}
