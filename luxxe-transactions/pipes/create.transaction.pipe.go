package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-transactions/messages"
	transactions_services "github.com/Emmanuella-codes/Luxxe/luxxe-transactions/services"
)

func CreateTransactionPipe(ctx context.Context, dto *dtos.CreateTransactionDTO) *shared.PipeRes[entities.Transaction] {
	transaction, err := transactions_services.CreateTransactionService(ctx, dto)
	if err != nil {
		return &shared.PipeRes[entities.Transaction]{
			Success: false,
			Message: messages.FAIL_CREATE_TRANSACTION,
		}
	}

	return &shared.PipeRes[entities.Transaction]{
		Success: true,
		Message: messages.SUCCESS_CREATE_TRANSACTION,
		Data:    transaction,
	}
}
