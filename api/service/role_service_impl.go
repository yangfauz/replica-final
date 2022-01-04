package service

import (
	"errors"
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/repository"
)

type roleServiceImpl struct {
	repository repository.RoleRepository
}

func NewRoleService(roleRepository *repository.RoleRepository) RoleService {
	return &roleServiceImpl{*roleRepository}
}

func (service *roleServiceImpl) GetAll() ([]entity.Role, error) {
	roles, err := service.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return roles, err
}

func (service *roleServiceImpl) GetById(id int) (entity.Role, error) {
	role, err := service.repository.FindById(id)

	if err != nil {
		return role, err
	}

	return role, nil
}

func (service *roleServiceImpl) Create(input model.CreateRoleRequest) (entity.Role, error) {
	//checkunique name
	roleName, err := service.repository.FindByName(input.Name)

	if err == nil {
		return roleName, errors.New("unique")
	}

	role := entity.Role{}
	role.Name = input.Name

	newRole, err := service.repository.Save(role)

	if err != nil {
		return newRole, err
	}

	return newRole, nil
}

func (service *roleServiceImpl) Update(id int, input model.CreateRoleRequest) (entity.Role, error) {
	//check id availability
	roleid, err := service.repository.FindById(id)

	if err != nil {
		return roleid, err
	}

	//checkuniquename
	roleName, err := service.repository.FindByName(input.Name)

	if err == nil {
		if roleid.Name != input.Name {
			return roleName, errors.New("unique")
		}
	}

	role := entity.Role{}
	role.ID = roleid.ID
	role.Name = input.Name

	updateRole, err := service.repository.Update(id, role)

	if err != nil {
		return updateRole, err
	}

	return updateRole, nil
}

func (service *roleServiceImpl) Delete(id int) bool {
	//check id availability
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
