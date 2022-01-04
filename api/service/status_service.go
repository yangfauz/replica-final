package service

import "replica-finalproject/api/entity"

type StatusService interface {
	GetByRole(roleId int) ([]entity.Status, error)
}
