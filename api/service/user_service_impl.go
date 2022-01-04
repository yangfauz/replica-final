package service

import (
	"errors"
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/repository"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{*userRepository}
}

func (service *userServiceImpl) GetAll() ([]entity.User, error) {
	users, err := service.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, err
}

func (service *userServiceImpl) GetById(id int) (entity.User, error) {
	user, err := service.repository.FindById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (service *userServiceImpl) Create(input model.CreateUserRequest) (entity.User, error) {
	//checkunique name
	userName, err := service.repository.FindByUsername(input.Username)

	if err == nil {
		return userName, errors.New("unique")
	}

	user := entity.User{}
	user.Name = input.Name
	user.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleID = input.RoleID

	newUser, err := service.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *userServiceImpl) Update(id int, input model.CreateUserRequest) (entity.User, error) {
	//check availability
	userid, err := service.repository.FindById(id)

	if err != nil {
		return userid, err
	}

	//checkuniquename
	userName, err := service.repository.FindByUsername(input.Username)

	if err == nil {
		if userid.Username != input.Username {
			return userName, errors.New("unique")
		}
	}

	user := entity.User{}
	user.ID = userid.ID
	user.Name = input.Name
	user.Username = input.Username

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.RoleID = input.RoleID

	updateUser, err := service.repository.Update(id, user)

	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func (service *userServiceImpl) Delete(id int) bool {
	//check availability
	_, err := service.repository.FindById(id)

	if err != nil {
		return false
	}

	_, err = service.repository.Delete(id)

	if err != nil {
		return false
	}

	return true
}
