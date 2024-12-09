package services

import (
	"context"
	"time"

	auth_services "github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	transaction_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/transactions"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/misc"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTransactionService(ctx context.Context, dto *dtos.CreateTransactionDTO) (*entities.Transaction, error) {
	pmtTransactionCtx 							:= dto.PmtTransactionCtx
	pmtTransactionCtxIDStr 					:= dto.PmtTransactionCtxID
	pmtTransactionInitiatorCtx 			:= dto.PmtTransactionInitiatorCtx
	pmtTransactionInitiatorCtxIDStr := dto.PmtTransactionInitiatorCtxID
	paymentClient 									:= dto.PaymentClient
	amount                          := dto.Amount
	fee                             := dto.Fee
	meta                            := dto.Meta

	txRefPrefix := "LUX_"

	txRefSuffCaps := auth_services.GetRandomCharacters(4, auth_services.CapitalLettersAlphabet)
	txRefSuffSmall := auth_services.GetRandomCharacters(8, auth_services.SmallLettersAlphabet)
	txRefSuffNums := auth_services.GetRandomCharacters(3, auth_services.NumbersAlphabet)

	txRefSuffix := txRefSuffCaps + txRefSuffSmall + txRefSuffNums
	transactionReference := txRefPrefix + txRefSuffix

	pmtTransactionCtxID := misc.StringToObjectID(pmtTransactionCtxIDStr)
	pmtTransactionInitiatorID := misc.StringToObjectID(pmtTransactionInitiatorCtxIDStr)

	transaction := entities.Transaction{
		ID: 													primitive.NewObjectID(),
		PmtTransactionCtx: 						pmtTransactionCtx,
		PmtTransactionCtxID: 					pmtTransactionCtxID,
		PmtTransactionInitiatorCtx: 	pmtTransactionInitiatorCtx,
		PmtTransactionInitiatorCtxID: pmtTransactionInitiatorID,
		TransactionReference: 				transactionReference,
		PaymentClient: 								paymentClient,
		PmtTransactionStatus: 				entities.PmtTxStatusPending,
		Amount: 											amount,
		Fee: 													fee,
		Meta: 												meta,
		CreatedAt: 										time.Now(),
	}
	return transaction_repo.TransactionRepo.Create(ctx, &transaction)
}
