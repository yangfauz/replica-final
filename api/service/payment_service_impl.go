package service

import (
	"replica-finalproject/api/entity"
	"replica-finalproject/api/model"
	"replica-finalproject/api/repository"
	"replica-finalproject/api/responder"
	"replica-finalproject/util"
	"time"
)

type paymentServiceImpl struct {
	repository repository.PaymentRepository
}

func NewPaymentService(paymentRepository *repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{*paymentRepository}
}

func (service *paymentServiceImpl) GetAllPaginate(pagination responder.Pagination) (responder.Pagination, error) {
	payments, err := service.repository.FindAllPaginate(pagination)

	if err != nil {
		return payments, err
	}

	return payments, err
}

func (service *paymentServiceImpl) GetAll() ([]entity.Payment, error) {
	payments, err := service.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return payments, err
}

func (service *paymentServiceImpl) GetById(id int) (entity.Payment, error) {
	payment, err := service.repository.FindById(id)

	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (service *paymentServiceImpl) Create(input model.CreatePaymentRequest, unit_id uint) (entity.Payment, error) {
	payment := entity.Payment{}
	payment.UserID = unit_id
	payment.RequestBy = input.RequestBy
	payment.Necessity = input.Necessity

	t, _ := time.Parse("2006-01-02", input.PaymentDate)

	payment.PaymentDate = t
	payment.PaymentAmount = input.PaymentAmount
	payment.PaymentCalculate = util.ToTerbilangRp(input.PaymentAmount)
	payment.PaymentAccountName = input.PaymentAccountName
	payment.PaymentAccountNumber = input.PaymentAccountNumber
	payment.StatusID = input.StatusID
	payment.Reason = input.Reason
	payment.CreatedAt = input.CreatedAt
	payment.UpdatedAt = input.UpdatedAt

	newPayment, err := service.repository.Save(payment)

	if err != nil {
		return newPayment, err
	}

	return newPayment, nil
}

func (service *paymentServiceImpl) Update(id int, input model.CreatePaymentRequest) (entity.Payment, error) {
	//check id availability
	paymentid, err := service.repository.FindById(id)

	if err != nil {
		return paymentid, err
	}

	payment := entity.Payment{}
	payment.ID = paymentid.ID
	payment.RequestBy = input.RequestBy
	payment.Necessity = input.Necessity

	t, _ := time.Parse("2006-01-02", input.PaymentDate)

	payment.PaymentDate = t
	payment.PaymentAmount = input.PaymentAmount
	payment.PaymentCalculate = util.ToTerbilangRp(input.PaymentAmount)
	payment.PaymentAccountName = input.PaymentAccountName
	payment.PaymentAccountNumber = input.PaymentAccountNumber
	payment.StatusID = input.StatusID
	payment.Reason = input.Reason
	payment.CreatedAt = input.CreatedAt
	payment.UpdatedAt = input.UpdatedAt

	updatePayment, err := service.repository.Update(id, payment)

	if err != nil {
		return updatePayment, err
	}

	return updatePayment, nil
}
