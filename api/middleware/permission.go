package middleware

import (
	"net/http"
	"replica-finalproject/api/responder"
	"replica-finalproject/exception"
	"replica-finalproject/util"

	"github.com/gofiber/fiber/v2"
)

//Admin
func RolePermissionAdmin(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	user_id := claims.UserId
	if user_id != 1 {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

//Unit
func RolePermissionUnit(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	user_id := claims.UserId
	if user_id != 2 {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

//General Support
func RolePermissionGs(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	user_id := claims.UserId
	if user_id != 3 {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}

//Accounting
func RolePermissionAc(c *fiber.Ctx) error {
	// Get claims from JWT.
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

	user_id := claims.UserId
	if user_id != 4 {
		return c.Status(http.StatusForbidden).JSON(responder.ApiResponse{
			Code:    http.StatusForbidden,
			Message: "Something Wrong",
			Error:   exception.NewString("Forbidden"),
			Data:    nil,
		})
	}

	return c.Next()
}
