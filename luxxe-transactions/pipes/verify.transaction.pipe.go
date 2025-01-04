package pipes

import (
	"context"
	"strconv"
	"sync"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	transaction_repo "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/transactions"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/messages"
	transactions_services "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/services"
)

// Global map to hold mutexes for each transaction reference
var transactionLocks = struct {
	sync.RWMutex
	locks map[string]*sync.Mutex
}{
	locks: make(map[string]*sync.Mutex),
}

func getTransactionLock(transactionReference string) *sync.Mutex {
	transactionLocks.RLock()
	lock, exists := transactionLocks.locks[transactionReference]
	transactionLocks.RUnlock()

	if !exists {
		transactionLocks.Lock()
		defer transactionLocks.Unlock()

		// Double-check after acquiring the write lock
		if lock, exists = transactionLocks.locks[transactionReference]; !exists {
			lock = &sync.Mutex{}
			transactionLocks.locks[transactionReference] = lock
		}
	}

	return lock
}

func VerifyTransactionPipe(ctx context.Context, dto *dtos.VerifyTransactionDTO) *shared.PipeRes[entities.Transaction] {
	transactionReference := dto.TransactionReference
	lock := getTransactionLock(transactionReference)

	// Lock the transaction-specific mutex
	lock.Lock()
	defer lock.Unlock()

	transaction, err := transaction_repo.TransactionRepo.QueryByTransactionReference(ctx, transactionReference)
	if err != nil {
		return &shared.PipeRes[entities.Transaction]{
			Success:  false,
			HookData: "",
			Message:  messages.NOT_FOUND_TRANSACTION,
		}
	}

	if transaction.PmtTransactionStatus == entities.PmtTxStatusSuccess {
		return &shared.PipeRes[entities.Transaction]{
			Success:  true,
			HookData: "transaction successful",
			Message:  messages.SUCCESS_VERIFY_TRANSACTION,
			Data:     transaction,
		}
	}

	if transaction.PmtTransactionStatus == entities.PmtTxStatusFailed {
		return &shared.PipeRes[entities.Transaction]{
			Success:  true,
			HookData: "transaction failed",
			Message:  messages.SUCCESS_VERIFY_TRANSACTION,
		}
	}

	success, pending, paymentClientVerificationResponse := transactions_services.PaymentClientVerifyTransaction(ctx, dto.PaymentClient, transaction, dto.PaymentClientFetchSingleResponse)
	if transaction.PaymentClientReference == "" {
		paymentClientReference := dto.PaymentClientReference
		if paymentClientReference == "" {
			switch dto.PaymentClient {
			case entities.PmtClientPaystack: 
			{
				paymentClientReference = strconv.Itoa(paymentClientVerificationResponse.Paystack.Data.ID)
			}
			}
		}
		transaction, _ = transaction_repo.TransactionRepo.UpdatePaymentClientReference(
			ctx,
			transaction.ID.Hex(),
			paymentClientReference,
		)
	}

	var updatedTransaction *entities.Transaction
	var updateErr error

	switch {
	case success:
			updatedTransaction, updateErr = transaction_repo.TransactionRepo.UpdateTransactionStatus(
					ctx, 
					transaction.ID.Hex(), 
					entities.PmtTxStatusSuccess,
			)
	case pending:
			updatedTransaction, updateErr = transaction_repo.TransactionRepo.UpdateTransactionStatus(
					ctx, 
					transaction.ID.Hex(), 
					entities.PmtTxStatusPending,
			)
	default: // failed
			updatedTransaction, updateErr = transaction_repo.TransactionRepo.UpdateTransactionStatus(
					ctx, 
					transaction.ID.Hex(), 
					entities.PmtTxStatusFailed,
			)
	}

	if updateErr != nil {
		return &shared.PipeRes[entities.Transaction]{
			Success:  false,
			HookData: "failed to update transaction status",
			Message:  messages.FAIL_UPDATE_TRANSACTION_STATUS,
		}
	}

	return &shared.PipeRes[entities.Transaction]{
		Success:   true,
		HookData: "transaction verified",
		Message:  messages.SUCCESS_VERIFY_TRANSACTION,
		Data:     updatedTransaction,
}
}
