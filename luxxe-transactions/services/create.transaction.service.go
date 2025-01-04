package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	auth_services "github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	transaction_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/transactions"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/misc"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTransactionService(ctx context.Context, dto *dtos.CreateTransactionDTO) (*entities.Transaction, error) {
	if dto == nil {
		return nil, errors.New("transaction DTO cannot be nil")
	}

	if dto.PmtTransactionCtx == "" || dto.PmtTransactionCtxID == "" || dto.PmtTransactionInitiatorCtxID == "" || dto.PaymentClient == "" {
		return nil, fmt.Errorf("missing required fields in transaction DTO")
	}

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

	pmtTransactionCtxID, err := misc.StringToObjectID(pmtTransactionCtxIDStr)
	if err != nil {
    return nil, fmt.Errorf("invalid PmtTransactionCtxID: %v", err)
	}
	pmtTransactionInitiatorID, err := misc.StringToObjectID(pmtTransactionInitiatorCtxIDStr)
	if err != nil {
    return nil, fmt.Errorf("invalid PmtTransactionInitiatorCtxID: %v", err)
}

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

	if transaction_repo.TransactionRepo == nil {
		return nil, fmt.Errorf("transaction repository is nil")
	}

	return transaction_repo.TransactionRepo.Create(ctx, &transaction)
}
