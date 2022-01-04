package controller

import (
	"net/http"
	"replica-finalproject/api/middleware"
	"replica-finalproject/api/model"
	"replica-finalproject/api/responder"
	"replica-finalproject/api/service"
	"replica-finalproject/exception"
	"replica-finalproject/util"

	"github.com/gofiber/fiber/v2"
)

type StatusController struct {
	StatusService service.StatusService
}

func NewStatusController(statusService *service.StatusService) StatusController {
	return StatusController{*statusService}
}

func (handler *StatusController) Route(app *fiber.App) {
	app.Get("/status", middleware.JWTProtected(), handler.StatusByRole)
}

func (handler *StatusController) StatusByRole(c *fiber.Ctx) error {
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
	role_id := claims.Role

	responses, err := handler.StatusService.GetByRole(int(role_id))
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
		Data:    model.FormatGetAllStatusResponse(responses),
	})

}
