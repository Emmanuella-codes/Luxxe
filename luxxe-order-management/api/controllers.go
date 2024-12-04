package api

import (
	"github.com/gofiber/fiber/v2"

	auth_messages "github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	cart_messages "github.com/Emmanuella-codes/Luxxe/luxxe-cart/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-order-management/pipes"
	cart_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/cart"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
)

func createOrder(ctx *fiber.Ctx) error {
  CreateOrder := new(dtos.CreateOrderDTO)

  if err := ctx.BodyParser(CreateOrder); err != nil {
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
	CreateOrder.UserID = userID

  cart, err := cart_repo.CartRepo.QueryByUserID(ctx.Context(), userID)
  if err != nil {
		statusCode = fiber.StatusBadRequest
		return ctx.Status(statusCode).JSON(
			fiber.Map{
				"statusCode": statusCode,
				"message":    cart_messages.NOT_FOUND_CART,
			},
		)
	}
	CreateOrder.CartID = cart.ID.Hex()

  success, err := shared_api.ValidateAPIData(CreateOrder)
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

  res := pipes.CreateOrderPipe(ctx.Context(), CreateOrder)
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

func updateOrder(ctx *fiber.Ctx) error {
  UpdateOrder := new(dtos.UpdateOrderDTO)

  if err := ctx.BodyParser(UpdateOrder); err != nil {
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
  UpdateOrder.UserID = userID

  success, err := shared_api.ValidateAPIData(UpdateOrder)
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
  res := pipes.UpdateOrderPipe(ctx.Context(), UpdateOrder)
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

func getOrder(ctx *fiber.Ctx) error {
  GetOrder := new(dtos.GetOrderDTO)

  if err := ctx.QueryParser(GetOrder); err != nil {
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
	GetOrder.UserID = userID

  success, err := shared_api.ValidateAPIData(GetOrder)
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

  res := pipes.GetOrderPipe(ctx.Context(), GetOrder)
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

func cancelOrder(ctx *fiber.Ctx) error {
  CancelOrder := new(dtos.CancelOrderDTO)

  if err := ctx.BodyParser(CancelOrder); err != nil {
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
  CancelOrder.UserID = userID

  success, err := shared_api.ValidateAPIData(CancelOrder)
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

  res := pipes.CancelOrderPipe(ctx.Context(), CancelOrder)
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
