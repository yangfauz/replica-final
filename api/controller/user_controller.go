package controller

import (
	"net/http"
	"replica-finalproject/api/middleware"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"
	"replica-finalproject/api/service"
	"replica-finalproject/api/validation"
	"replica-finalproject/exception"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{*userService}
}

func (handler *UserController) Route(app *fiber.App) {
	app.Get("/users", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetAllUser)
	app.Get("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetUserById)
	app.Post("/users", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.CreateUser)
	app.Put("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.UpdateUser)
	app.Delete("/users/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.DeleteUser)
}

func (handler *UserController) GetAllUser(c *fiber.Ctx) error {
	responses, err := handler.UserService.GetAll()
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
		Data:    model.FormatGetAllUserResponse(responses),
	})
}

func (handler *UserController) GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.UserService.GetById(id)

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
		Data:    model.FormatGetUserResponse(responses),
	})
}

func (handler *UserController) CreateUser(c *fiber.Ctx) error {
	var input model.CreateUserRequest

	err := c.BodyParser(&input)

	if err != nil {
		return err
	}

	//validate input
	validation.UserValidate(input)

	responses, err := handler.UserService.Create(input)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusUnprocessableEntity).JSON(responder.ApiResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "Something Wrong",
				Error:   exception.NewString("username: Username Available"),
				Data:    nil,
			})
		}
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
		Data:    model.FormatCreateUserResponse(responses),
	})
}

func (handler *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.CreateUserRequest
	err = c.BodyParser(&input)

	//validate input
	validation.UserValidate(input)

	if err != nil {
		//error
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.UserService.Update(id, input)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusUnprocessableEntity).JSON(responder.ApiResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "Something Wrong",
				Error:   exception.NewString("username: Username Available"),
				Data:    nil,
			})
		}
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
		Data:    model.FormatCreateUserResponse(responses),
	})
}

func (handler *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses := handler.UserService.Delete(id)

	if responses != true {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete Data Failed",
			Error:   exception.NewString("Delete Failed / Record Not Found"),
			Data:    false,
		})
	}
	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Delete Data Success",
		Error:   nil,
		Data:    responses,
	})
}
