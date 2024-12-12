package services

import (
	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	transaction_types "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/types"
)

func InitializePaystackTransactionService(
	transaction *entities.Transaction, 
	order *entities.OrderManagement,
	) (*transaction_types.PaystackInitializeTransactionResponseData, error) {
		secretKey := config.EnvConfig.PAYSTACK_TEST_SECRET_KEY

		initializeReq := map[string]interface{}{
			"amount":    order.CartTotal * 100,
			"email":     order.Email,
			"reference": transaction.TransactionReference,
		}

		response, err := shared.POST[transaction_types.PaystackInitializeTransactionResponseData](
			shared.HttpUtilsReq{
				BaseRoute: config.EnvConfig.PAYSTACK_API,
				Ext:       "/transaction/initialize",
				Token:     secretKey,
				Body:      initializeReq,
			},
		)
	
		if err != nil {
			return nil, err
	}

	return &response, nil
}
