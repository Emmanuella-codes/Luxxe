package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/pipes"
	auth_services "github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
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

func sendOtp(ctx *fiber.Ctx) error {
	SendOTP := new(dtos.SendOTPDTO)

	if err := ctx.BodyParser(SendOTP); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(SendOTP)
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
	res := pipes.SendOTPPipe(ctx.Context(), SendOTP)
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
		},
	)
}

func resetUserPasswordByEmail(ctx *fiber.Ctx) error {
	ResetUserPassword := new(dtos.ResetUserPasswordByEmailDTO)

	if err := ctx.BodyParser(ResetUserPassword); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(ResetUserPassword)
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
	res := pipes.ResetUserPasswordByEmailPipe(ctx.Context(), ResetUserPassword)
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
		},
	)
}

func resetUserPasswordByID(ctx *fiber.Ctx) error {
	ResetUserPasswordByUserID := new(dtos.ResetUserPasswordByUserIDDTO)

	if err := ctx.BodyParser(ResetUserPasswordByUserID); err != nil {
		return err
	}

	var statusCode int
	AccountToken := ctx.Locals("token").(*auth_services.AccountTokenStruct)
	userID := AccountToken.UserID

	user, err := repo_user.UserRepo.QueryByID(ctx.Context(), userID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    messages.NOT_FOUND_USER,
			},
		)
	}

	ResetUserPasswordByUserID.UserID = userID
	success, err := shared_api.ValidateAPIData(ResetUserPasswordByUserID)
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

	res := pipes.ResetUserPasswordByEmailPipe(ctx.Context(), 
		&dtos.ResetUserPasswordByEmailDTO{
			Email: 		user.Email,
			Otp:   		ResetUserPasswordByUserID.Otp,
			Password: ResetUserPasswordByUserID.Password,
		},
	)

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
		},
	)
}

func verifyEmail(ctx *fiber.Ctx) error {
	VerifyEmail := new(dtos.VerifyEmailDTO)

	if err := ctx.QueryParser(VerifyEmail); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(VerifyEmail)
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

	VerifyEmail.UserID = ctx.Query("userID")

	pipes.VerifyEmailPipe(ctx.Context(), VerifyEmail)

	settingsRoute := config.EnvConfig.FRONTEND_APP_URL + "/settings?email-verified=set"

	return ctx.Redirect(settingsRoute)
}
