package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type GetAllInitiatorsTransactionsDTO struct {
	Page                      	 int         							 	`query:"page" validate:"required"`
	PmtTransactionCtx            entities.PmtTransactionCtx `query:"pmtTransactionCtx" validate:"required,oneof=order"`
	PmtTransactionInitiatorCtxID string      							 	`query:"pmtTransactionInitiatorCtxID" validate:"required,min=24,max=24"`
}
