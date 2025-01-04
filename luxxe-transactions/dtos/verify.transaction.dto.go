package dtos

import (
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	transaction_types "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/types"
)

type VerifyTransactionDTO struct {
	PaymentClient                    entities.PaymentClient                             `json:"paymentClient" validate:"required,oneof=paystack"`
	PaymentClientReference           string                                             `json:"paymentClientReference"`
	TransactionReference             string                                             `json:"transactionReference" validate:"required"`
	PaymentClientFetchSingleResponse transaction_types.PaymentClientFetchSingleResponse `json:"paymentClientFetchSingleResponse"`
}
