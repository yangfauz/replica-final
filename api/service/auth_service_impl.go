package service

import (
	"errors"
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/repository"

	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repository repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) AuthService {
	return &authServiceImpl{*userRepository}
}

func (service *authServiceImpl) Login(input model.UserLoginRequest) (entity.User, error) {
	username := input.Username
	password := input.Password

	user, err := service.repository.FindByUsername(username)

	if err != nil {
		return user, errors.New("Username Not Found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Username/Password was wrong")
	}

	return user, nil
}
