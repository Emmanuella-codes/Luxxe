package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"

	"github.com/gofiber/fiber/v2"

	auth_messages "github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	auth_services "github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/pipes"
	transaction_types "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/types"
)

func verifyTransaction(ctx *fiber.Ctx) error {
	VerifyTransaction := new(dtos.VerifyTransactionDTO)

	if err := ctx.BodyParser(VerifyTransaction); err != nil {
		return err
	}

	success, err := shared_api.ValidateAPIData(VerifyTransaction)
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
	res := pipes.VerifyTransactionPipe(ctx.Context(), VerifyTransaction)
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

func paystackWebHook(ctx *fiber.Ctx) error {
	requestBody := ctx.Body()
	secretKey := config.EnvConfig.PAYSTACK_TEST_SECRET_KEY

	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write(requestBody)

	hash := hex.EncodeToString(h.Sum(nil))
	paystackSignature := ctx.Get("x-paystack-signature")

	if hash == paystackSignature {
		PaystackStandardWebhookVerificationResponse := new(transaction_types.PaystackStandardWebhookVerificationResponse)
		if err := ctx.BodyParser(PaystackStandardWebhookVerificationResponse); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
		}

		if PaystackStandardWebhookVerificationResponse.Event == transaction_types.WebHookEventPaystackSuccessfulCharge {
			res := PaystackStandardWebhookVerificationResponse.Data
			paymentClientReference := strconv.Itoa(res.ID)
			transactionReference := res.Reference
			paymentClientFetchSingleResponse := transaction_types.PaymentClientFetchSingleResponse{}
			paymentClientFetchSingleResponse.Paystack = res

			verifyTransactionDTO := &dtos.VerifyTransactionDTO{
				PaymentClient: entities.PmtClientPaystack,
				PaymentClientReference: paymentClientReference,
				TransactionReference:  transactionReference,
				PaymentClientFetchSingleResponse: paymentClientFetchSingleResponse,
			}
			go pipes.VerifyTransactionPipe(ctx.Context(), verifyTransactionDTO)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func getAllInitiatorsTransactions(ctx *fiber.Ctx) error {
	GetAllInitiatorsTransactions := new(dtos.GetAllInitiatorsTransactionsDTO)

	if err := ctx.QueryParser(GetAllInitiatorsTransactions); err != nil {
		return err
	}

	AccountToken := ctx.Locals("token").(*auth_services.AccountTokenStruct)
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

	GetAllInitiatorsTransactions.PmtTransactionInitiatorCtxID = userID
	success, err := shared_api.ValidateAPIData(GetAllInitiatorsTransactions)
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

	res := pipes.GetAllInitiatorsTransactionsPipe(ctx.Context(), GetAllInitiatorsTransactions)
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
