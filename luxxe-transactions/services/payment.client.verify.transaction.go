package services

import (
	"context"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	transaction_types "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/types"
)


type PaymentClientVerificationResponse struct {
	Paystack transaction_types.PaystackResponse[transaction_types.PaystackStandardVerificationResponseData]
}

func PaymentClientVerifyTransaction(ctx context.Context,
	paymentClient entities.PaymentClient,
	transaction *entities.Transaction,
	paymentClientFetchSingleResponse ...transaction_types.PaymentClientFetchSingleResponse,
) (bool, bool, PaymentClientVerificationResponse) {
	var success bool
	var pending bool
	paymentClientResponse := PaymentClientVerificationResponse{}

	if paymentClient == entities.PmtClientPaystack {
		pstkSuccess, pstkPending, pstkRes := verifyTransactionPaystack(transaction.PaymentClientReference, paymentClientFetchSingleResponse[0])
		success = pstkSuccess
		pending = pstkPending
		paymentClientResponse.Paystack = pstkRes
	}
	return success, pending, paymentClientResponse
}

func verifyTransactionPaystack(
	transactionReference string,
	paymentClientFetchSingleResponse ...transaction_types.PaymentClientFetchSingleResponse,
) (bool, bool, transaction_types.PaystackResponse[transaction_types.PaystackStandardVerificationResponseData]) {
	secretKey := config.EnvConfig.PAYSTACK_TEST_SECRET_KEY
	if config.EnvConfig.APP_ENV == "prodction" {
		secretKey = config.EnvConfig.PAYSTACK_LIVE_SECRET_KEY
	}

	res := transaction_types.PaystackResponse[transaction_types.PaystackStandardVerificationResponseData]{}

	useSingleResponse := len(paymentClientFetchSingleResponse) > 0
	if useSingleResponse {
		res.Data = paymentClientFetchSingleResponse[0].Paystack
	} else {
		res, err := shared.GET[transaction_types.PaystackResponse[transaction_types.PaystackStandardVerificationResponseData]](
			shared.HttpUtilsReq{
				BaseRoute: config.EnvConfig.PAYSTACK_API,
				Ext:       "/transaction/verify/" + transactionReference,
				Token:     secretKey,
			},
		)
		if err != nil {
			return false, true, res
		}
	}

	if res.Data.Status == transaction_types.TransactionStatusPaystackSuccess {
		return true, false, res
	}

	if res.Data.Status == transaction_types.TransactionStatusPaystackFailed {
		return false, true, res
	}

	return true, false, res
}
