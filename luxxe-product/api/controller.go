package api

import (
	"github.com/gofiber/fiber/v2"

	auth_messages "github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/dtos"
	product_messages "github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-product/pipes"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
)

func createProduct(ctx *fiber.Ctx) error {
	CreateProduct := new(dtos.CreateProductDTO)

	if err := ctx.BodyParser(CreateProduct); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	var statusCode int
	user, err := repo_user.UserRepo.QueryByID(ctx.Context(), userID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    auth_messages.NOT_FOUND_USER,
			},
		)
	}

	if user.AccountRole != entities.AccountRoleAdmin {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    product_messages.NOT_AUTHORIZED,
			},
		)
	}
	CreateProduct.UserID = userID

	success, err := shared_api.ValidateAPIData(CreateProduct)
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

	res := pipes.CreateProductPipe(ctx.Context(), CreateProduct)
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

func updateProduct(ctx *fiber.Ctx) error {
	UpdateProduct := new(dtos.UpdateProductDTO)

	if err := ctx.BodyParser(UpdateProduct); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	var statusCode int
	user, err := repo_user.UserRepo.QueryByID(ctx.Context(), userID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    auth_messages.NOT_FOUND_USER,
			},
		)
	}

	if user.AccountRole != entities.AccountRoleAdmin {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    product_messages.NOT_AUTHORIZED,
			},
		)
	}

	success, err := shared_api.ValidateAPIData(UpdateProduct)
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

	res := pipes.UpdateProductPipe(ctx.Context(), UpdateProduct)
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

func getProducts(ctx *fiber.Ctx) error {
	GetProducts := new(dtos.GetProductDTO)

	if err := ctx.QueryParser(GetProducts); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	var statusCode int
	_, err := repo_user.UserRepo.QueryByID(ctx.Context(), userID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    auth_messages.NOT_FOUND_USER,
			},
		)
	}
	
	success, err := shared_api.ValidateAPIData(GetProducts)
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

	res := pipes.GetAllProductsPipe(ctx.Context(), GetProducts)
	if res.Success {
		statusCode = fiber.StatusOK
	} else {
		statusCode = fiber.StatusBadRequest
	}

	var payload interface{}
	if res.Data == nil {
		payload = []interface{}{} // Return an empty array if data is nil or empty string
	} else {
		payload = res.Data
	}

	return ctx.Status(statusCode).JSON(
		fiber.Map{
			"statusCode": statusCode,
			"message":    res.Message,
			"payload":    payload,
			"token":      res.Token,
		},
	)
}

func deleteProduct(ctx *fiber.Ctx) error {
	DeleteProduct := new(dtos.DeleteProductDTO)

	if err := ctx.BodyParser(DeleteProduct); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	var statusCode int
	user, err := repo_user.UserRepo.QueryByID(ctx.Context(), userID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    auth_messages.NOT_FOUND_USER,
			},
		)
	}

	if user.AccountRole != entities.AccountRoleAdmin {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    product_messages.NOT_AUTHORIZED,
			},
		)
	}

	success, err := shared_api.ValidateAPIData(DeleteProduct)
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

	res := pipes.DeleteProductPipe(ctx.Context(), DeleteProduct)
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
