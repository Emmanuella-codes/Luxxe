package transactions

import (
	"context"

	"github.com/go-kit/log"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error)
	QueryPendingTransactionByTransactionReference(ctx context.Context, transactionReference string) (*entities.Transaction, error)
	QueryAllTransactionsByTransactionInitiatorCtxID(
		ctx context.Context, 
		transactionCtx entities.PmtTransactionCtx, 
		transactionInitiatorCtxID string, page int,
	) (*[]entities.Transaction, int64, error)
	QueryTransactionsByOrderID(
		ctx context.Context,
		orderID string,
		page int,
		PmtTransactionStatus entities.PmtTransactionStatus,
	) (*[]entities.Transaction, int64, error)
	QueryByTransactionReference(ctx context.Context, transactionReference string) (*entities.Transaction, error)
	UpdateTransactionStatus(ctx context.Context, txID string, txStatus entities.PmtTransactionStatus) (*entities.Transaction, error)
	UpdatePaymentClientReference(ctx context.Context, txID string, paymentClientReference string) (*entities.Transaction, error)
}

var TransactionRepo TransactionRepository

func InitTransactionRepo(logger *log.Logger) {
	TransactionRepo = newMgRepository(logger)
}
