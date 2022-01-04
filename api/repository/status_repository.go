package repository

import "replica-finalproject/api/entity"

type StatusRepository interface {
	FindByRole(statusId []int) ([]entity.Status, error)
}
