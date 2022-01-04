package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"
)

type PaymentService interface {
	GetAllPaginate(pagination responder.Pagination) (responder.Pagination, error)
	GetAll() ([]entity.Payment, error)
	GetById(id int) (entity.Payment, error)
	Create(input model.CreatePaymentRequest, unit_id uint) (entity.Payment, error)
	Update(id int, input model.CreatePaymentRequest) (entity.Payment, error)
}
