package repository

import (
	"replica-finalproject/api/entity"

	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	database *gorm.DB
}

func NewRoleRepository(database *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{database}
}

func (repository *roleRepositoryImpl) FindAll() ([]entity.Role, error) {
	var roles []entity.Role

	err := repository.database.Find(&roles).Error
	if err != nil {
		return roles, err
	}

	return roles, nil
}

func (repository *roleRepositoryImpl) FindById(id int) (entity.Role, error) {
	var role entity.Role

	err := repository.database.First(&role, id).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repository *roleRepositoryImpl) FindByName(name string) (entity.Role, error) {
	var role entity.Role

	err := repository.database.Where("name = ?", name).First(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repository *roleRepositoryImpl) Save(role entity.Role) (entity.Role, error) {
	err := repository.database.Create(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repository *roleRepositoryImpl) Update(id int, role entity.Role) (entity.Role, error) {
	err := repository.database.Where("id = ?", id).Updates(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}

func (repository *roleRepositoryImpl) Delete(id int) (entity.Role, error) {
	var role entity.Role

	err := repository.database.Where("id = ?", id).Delete(&role).Error

	if err != nil {
		return role, err
	}

	return role, nil
}
