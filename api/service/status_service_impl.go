package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/repository"
)

type statusServiceImpl struct {
	repository repository.StatusRepository
}

func NewStatusService(statusRepository *repository.StatusRepository) StatusService {
	return &statusServiceImpl{*statusRepository}
}

func (service *statusServiceImpl) GetByRole(roleId int) ([]entity.Status, error) {
	//bandingkan role nya
	var statusId []int

	if roleId == 3 {
		statusId = []int{2, 3}
	} else if roleId == 4 {
		statusId = []int{4, 5}
	} else {
		statusId = []int{}
	}

	status, err := service.repository.FindByRole(statusId)

	if err != nil {
		return status, err
	}

	return status, nil
}
