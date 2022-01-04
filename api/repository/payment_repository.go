package repository

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/responder"
)

type PaymentRepository interface {
	FindAllPaginate(pagination responder.Pagination) (responder.Pagination, error)
	FindAll() ([]entity.Payment, error)
	FindById(id int) (entity.Payment, error)
	Save(payment entity.Payment) (entity.Payment, error)
	Update(id int, payment entity.Payment) (entity.Payment, error)
}
