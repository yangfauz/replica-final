package repository

import (
	"replica-finalproject/api/entity"

	"gorm.io/gorm"
)

type statusRepositoryImpl struct {
	database *gorm.DB
}

func NewStatusRepository(database *gorm.DB) StatusRepository {
	return &statusRepositoryImpl{database}
}

func (repository *statusRepositoryImpl) FindByRole(statusId []int) ([]entity.Status, error) {
	var statuses []entity.Status

	err := repository.database.Find(&statuses, statusId).Error
	if err != nil {
		return statuses, err
	}

	return statuses, nil
}
