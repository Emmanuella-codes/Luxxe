package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	transaction_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/transactions"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/messages"
)

type Transactions struct {
	TransactionCount int64                   `json:"transactionCount"`
	TransactionList  *[]entities.Transaction `json:"transactionList"`
}

func GetAllInitiatorsTransactionsPipe(ctx context.Context, dto *dtos.GetAllInitiatorsTransactionsDTO) *shared.PipeRes[Transactions] {
	pmtTransactionCtx, page, pmtTransactionInitiatorCtxID := dto.PmtTransactionCtx, dto.Page, dto.PmtTransactionInitiatorCtxID

	transactionList, transactionsCount, _ := transaction_repo.TransactionRepo.QueryAllTransactionsByTransactionInitiatorCtxID(ctx, pmtTransactionCtx, pmtTransactionInitiatorCtxID, page)
	
	return &shared.PipeRes[Transactions]{
		Success: true,
		Message: messages.SUCCESS_GET_TRANSACTION_LIST,
		Data: &Transactions{
			TransactionCount: transactionsCount,
			TransactionList: transactionList,
		},
	} 
}