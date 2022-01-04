package exception

import (
	"net/http"
	"replica-finalproject/api/responder"

	"github.com/gofiber/fiber/v2"
)

func NewString(s string) *string {
	return &s
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)

	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(responder.ApiResponse{
			// Status: "BAD_REQUEST",
			Code:    http.StatusBadRequest,
			Message: "Something Wrong",
			Error:   NewString(err.Error()),
			Data:    nil,
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(responder.ApiResponse{
		// Status: "INTERNAL_SERVER_ERROR",
		Code:    http.StatusInternalServerError,
		Message: "Something Wrong",
		Error:   NewString(err.Error()),
		Data:    nil,
	})
}
