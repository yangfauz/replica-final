package controller

import (
	"net/http"
	"replica-finalproject/api/middleware"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"
	"replica-finalproject/api/service"
	"replica-finalproject/api/validation"
	"replica-finalproject/exception"
	"replica-finalproject/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) PaymentController {
	return PaymentController{*paymentService}
}

func (handler *PaymentController) Route(app *fiber.App) {
	app.Get("/paginate-payments", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetAllPaginatePayment)
	app.Get("/payments", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetAllPayment)
	app.Get("/payments/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetPaymentById)
	app.Post("/payments", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.CreatePayment)
	app.Put("/payments/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.UpdatePayment)
}

func (handler *PaymentController) GetAllPaginatePayment(c *fiber.Ctx) error {
	//body
	limit, err := strconv.Atoi(c.FormValue("limit"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("limit required."),
			Data:    nil,
		})
	}

	page, err := strconv.Atoi(c.FormValue("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   exception.NewString("page required."),
			Data:    nil,
		})
	}

	keyword := c.FormValue("keyword")

	set_paginate := responder.Pagination{}
	set_paginate.Limit = limit
	set_paginate.Page = page
	set_paginate.Keyword = keyword
	set_paginate.Sort = "Id asc"

	responses, err := handler.PaymentService.GetAllPaginate(set_paginate)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    responder.Pagination(responses),
	})
}

func (handler *PaymentController) GetAllPayment(c *fiber.Ctx) error {
	responses, err := handler.PaymentService.GetAll()
	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    model.FormatGetAllPaymentResponse(responses),
	})
}

func (handler *PaymentController) GetPaymentById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.PaymentService.GetById(id)

	if err != nil {
		//error
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Get Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Get Data Success",
		Error:   nil,
		Data:    model.FormatGetPaymentResponse(responses),
	})
}

func (handler *PaymentController) CreatePayment(c *fiber.Ctx) error {
	var input model.CreatePaymentRequest

	err := c.BodyParser(&input)

	if err != nil {
		return err
	}

	//validate input
	validation.PaymentValidate(input)

	//claim token
	claims, err := util.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}
	unit_id := claims.UserId

	responses, err := handler.PaymentService.Create(input, uint(unit_id))

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Create Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(responder.ApiResponse{
		Code:    http.StatusCreated,
		Message: "Create Data Success",
		Error:   nil,
		Data:    model.FormatCreatePaymentResponse(responses),
	})
}

func (handler *PaymentController) UpdatePayment(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.CreatePaymentRequest
	err = c.BodyParser(&input)

	//validate input
	validation.PaymentValidate(input)

	if err != nil {
		//error
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.PaymentService.Update(id, input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Update Data Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Update Data Success",
		Error:   nil,
		Data:    model.FormatCreatePaymentResponse(responses),
	})
}
