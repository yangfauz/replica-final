package controller

import (
	"net/http"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"
	"replica-finalproject/api/service"
	"replica-finalproject/api/validation"
	"replica-finalproject/exception"
	"replica-finalproject/util"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService *service.AuthService) AuthController {
	return AuthController{*authService}
}

func (handler *AuthController) Route(app *fiber.App) {
	app.Post("/login", handler.Login)
}

func (handler *AuthController) Login(c *fiber.Ctx) error {
	var input model.UserLoginRequest

	err := c.BodyParser(&input)

	if err != nil {
		return err
	}

	//validate input
	validation.LoginValidate(input)

	responses, err := handler.AuthService.Login(input)

	if err != nil {
		//error
		return c.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	token, err := util.GenerateNewAccessToken(responses)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: "Login Failed",
			Error:   exception.NewString(err.Error()),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(responder.ApiResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Error:   nil,
		Data:    model.FormatLoginUserResponse(responses, token),
	})
}
