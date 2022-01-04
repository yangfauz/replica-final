package repository

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"

	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	database *gorm.DB
}

func NewPaymentRepository(database *gorm.DB) PaymentRepository {
	return &paymentRepositoryImpl{database}
}

func (repository *paymentRepositoryImpl) FindAllPaginate(pagination responder.Pagination) (responder.Pagination, error) {
	var payments []entity.Payment

	keyword := "%" + pagination.Keyword + "%"
	err := repository.database.
		Where("request_by LIKE ?", keyword).
		Or("payment_account_name LIKE ?", keyword).
		Or("payment_account_number LIKE ?", keyword).
		Scopes(responder.Paginate(keyword, payments, &pagination, repository.database)).
		Find(&payments).Error

	if err != nil {
		return pagination, err
	}

	pagination.Rows = model.FormatGetAllPaymentResponse(payments)

	return pagination, nil
}

func (repository *paymentRepositoryImpl) FindAll() ([]entity.Payment, error) {
	var payment []entity.Payment

	err := repository.database.Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (repository *paymentRepositoryImpl) FindById(id int) (entity.Payment, error) {
	var payment entity.Payment

	err := repository.database.First(&payment, id).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (repository *paymentRepositoryImpl) Save(payment entity.Payment) (entity.Payment, error) {
	err := repository.database.Create(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (repository *paymentRepositoryImpl) Update(id int, payment entity.Payment) (entity.Payment, error) {
	err := repository.database.Where("id = ?", id).Updates(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}
