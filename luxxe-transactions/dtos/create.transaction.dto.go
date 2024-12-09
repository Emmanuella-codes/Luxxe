package dtos

import entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"

type CreateTransactionDTO struct {
	PmtTransactionCtx 						entities.PmtTransactionCtx 					`json:"pmtTransactionCtx" validate:"required,oneof=order"`
	PmtTransactionCtxID 					string 															`json:"pmtTransactionCtxID" validate:"required,min=24,max=24"`
	PmtTransactionInitiatorCtx 		entities.PmtTransactionInitiatorCtx `json:"pmtTransactionInitiatorCtx" validate:"required,oneof=account"`
	PmtTransactionInitiatorCtxID 	string 															`json:"pmtTransactionInitiatorCtxID" validate:"required,min=24,max=24"`
	PaymentClient 								entities.PaymentClient 							`json:"paymentClient" validate:"required,oneof=paystack"`
	Meta                      		map[string]interface{}           		`json:"meta"`
	Amount                    		float64                          		`json:"amount" validate:"required"`
	Fee                       		float64                          		`json:"fee"`
}