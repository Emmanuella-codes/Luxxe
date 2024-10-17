package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	"github.com/Emmanuella-codes/Luxxe/luxxe-profile/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-profile/pipes"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
)

func getUserByAccountToken(ctx *fiber.Ctx) error {
	GetUser := new(dtos.GetUserProfileDTO)

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	GetUser.UserID = AccountToken.UserID

	success, err := shared_api.ValidateAPIData(GetUser)
	if !success {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": fiber.StatusBadRequest,
				"message":    "Invalid request data",
				"payload":    map[string]string{},
				"error":      err.Error(),
			},
		)
	}

	var statusCode int
	res := pipes.GetUserProfilePipe(ctx.Context(), GetUser)
	if res.Success {
		statusCode = fiber.StatusOK
	} else {
		statusCode = fiber.StatusBadRequest
	}

	return ctx.Status(statusCode).JSON(
		fiber.Map{
			"statusCode": statusCode,
			"message":    res.Message,
			"payload":    res.Data,
			"token":      res.Token,
		},
	)
}
