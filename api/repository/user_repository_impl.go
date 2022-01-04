package repository

import (
	"replica-finalproject/api/entity"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{database}
}

func (repository *userRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := repository.database.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (repository *userRepositoryImpl) FindById(id int) (entity.User, error) {
	var user entity.User

	err := repository.database.First(&user, id).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	err := repository.database.Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Save(user entity.User) (entity.User, error) {
	err := repository.database.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Update(id int, user entity.User) (entity.User, error) {
	err := repository.database.Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Delete(id int) (entity.User, error) {
	var user entity.User

	err := repository.database.Where("id = ?", id).Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
