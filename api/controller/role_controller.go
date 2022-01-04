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

type RoleController struct {
	RoleService service.RoleService
}

func NewRoleController(roleService *service.RoleService) RoleController {
	return RoleController{*roleService}
}

func (handler *RoleController) Route(app *fiber.App) {
	app.Get("/roles", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetAllRole)
	app.Get("/roles/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.GetRoleById)
	app.Post("/roles", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.CreateRole)
	app.Put("/roles/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.UpdateRole)
	app.Delete("/roles/:id", middleware.JWTProtected(), middleware.RolePermissionAdmin, handler.DeleteRole)
}

func (handler *RoleController) GetAllRole(c *fiber.Ctx) error {
	responses, err := handler.RoleService.GetAll()
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
		Data:    model.FormatGetAllRoleResponse(responses),
	})
}

func (handler *RoleController) GetRoleById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.RoleService.GetById(id)

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
		Data:    model.FormatGetRoleResponse(responses),
	})
}

func (handler *RoleController) CreateRole(c *fiber.Ctx) error {
	var input model.CreateRoleRequest

	err := c.BodyParser(&input)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	//validate input
	validation.RoleValidate(input)

	responses, err := handler.RoleService.Create(input)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusUnprocessableEntity).JSON(responder.ApiResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "Something Wrong",
				Error:   exception.NewString("name: Name Available"),
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
		Data:    model.FormatCreateRoleResponse(responses),
	})
}

func (handler *RoleController) UpdateRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	var input model.CreateRoleRequest
	err = c.BodyParser(&input)

	//validate input
	validation.RoleValidate(input)

	if err != nil {
		//error
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses, err := handler.RoleService.Update(id, input)

	if err != nil {
		//error
		if err.Error() == "unique" {
			return c.Status(http.StatusUnprocessableEntity).JSON(responder.ApiResponse{
				Code:    http.StatusUnprocessableEntity,
				Message: "Something Wrong",
				Error:   exception.NewString("name: Name Available"),
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
		Data:    model.FormatCreateRoleResponse(responses),
	})
}

func (handler *RoleController) DeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responder.ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Something Wrong",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	responses := handler.RoleService.Delete(id)

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
