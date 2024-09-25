package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/pipes"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
)

func signUpUser(ctx *fiber.Ctx) error {
	SignUpUser := new(dtos.SignUpUserDTO)

	if err := ctx.BodyParser(SignUpUser); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(SignUpUser)
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
	res := pipes.SignUpUserPipe(ctx.Context(), SignUpUser)
	if res.Success {
		statusCode = fiber.StatusCreated
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

func signInUser(ctx *fiber.Ctx) error {
	SignInUser := new(dtos.SignInUserDTO)

	if err := ctx.BodyParser(SignInUser); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(SignInUser)
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
	res := pipes.SignInUserPipe(ctx.Context(), SignInUser)
	if res.Success {
		statusCode = fiber.StatusCreated
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
