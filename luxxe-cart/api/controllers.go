package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	auth_messages "github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-cart/pipes"
	product_messages "github.com/Emmanuella-codes/Luxxe/luxxe-product/messages"
	product_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
)

func addToCart(ctx *fiber.Ctx) error {
	AddToCart := new(dtos.AddToCartDTO)

	if err := ctx.BodyParser(AddToCart); err != nil {
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
	AddToCart.UserID = userID

	product, err := product_repo.ProductRepo.QueryByID(ctx.Context(), AddToCart.ProductID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    product_messages.NOT_FOUND_PRODUCT,
			},
		)
	}
	AddToCart.Price = product.Price 

	success, err := shared_api.ValidateAPIData(AddToCart)
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

	res := pipes.AddToCartPipe(ctx.Context(), AddToCart)
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

func updateCart(ctx *fiber.Ctx) error {
	UpdateCart := new(dtos.UpdateCartItemDTO)

	if err := ctx.BodyParser(UpdateCart); err != nil {
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
	UpdateCart.UserID = userID

	product, err := product_repo.ProductRepo.QueryByID(ctx.Context(), UpdateCart.ProductID)
	if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    product_messages.NOT_FOUND_PRODUCT,
			},
		)
	}
	UpdateCart.Price = product.Price 

	success, err := shared_api.ValidateAPIData(UpdateCart)
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

	res := pipes.UpdateCartItemPipe(ctx.Context(), UpdateCart)
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

func getCart(ctx *fiber.Ctx) error {
	GetCart := new(dtos.GetCartDTO)

	if err := ctx.QueryParser(GetCart); err != nil {
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
	GetCart.UserID = userID
	
	success, err := shared_api.ValidateAPIData(GetCart)
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

	res := pipes.GetCartPipe(ctx.Context(), GetCart)
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

func removeItemFromCart(ctx *fiber.Ctx) error {
	RemoveItemFromCart := new(dtos.RemoveItemFromCartDTO)

	if err := ctx.BodyParser(RemoveItemFromCart); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	fmt.Println("UserID from AccountToken:", userID)

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
	RemoveItemFromCart.UserID = userID
	
	success, err := shared_api.ValidateAPIData(RemoveItemFromCart)
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

	res := pipes.RemoveItemFromCartPipe(ctx.Context(), RemoveItemFromCart)
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

func clearCart(ctx *fiber.Ctx) error {
	ClearCart := new(dtos.ClearCartDTO)

	// if err := ctx.BodyParser(ClearCart); err != nil {
	// 	return err
	// }

	AccountToken := ctx.Locals("token").(*services.AccountTokenStruct)
	userID := AccountToken.UserID

	fmt.Println("UserID from AccountToken:", userID)

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
	ClearCart.UserID = userID

	success, err := shared_api.ValidateAPIData(ClearCart)
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

	res := pipes.ClearCartPipe(ctx.Context(), ClearCart)
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
