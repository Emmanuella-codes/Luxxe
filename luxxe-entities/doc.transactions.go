package entities

import (
	"time"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PmtTransactionCtx string

const PmtTxCtxOrder PmtTransactionCtx = "order"

type PmtTransactionInitiatorCtx string

const PmtTxInitCtxAccount PmtTransactionInitiatorCtx = "account"

type PaymentClient string

const PmtClientPaystack PaymentClient = "paystack"

type PmtTransactionStatus string

const PmtTxStatusSuccess PmtTransactionStatus = "success"
const PmtTxStatusPending PmtTransactionStatus = "pending"
const PmtTxStatusFailed  PmtTransactionStatus = "failed"

type Transaction struct {
	ID                        	 primitive.ObjectID      		`json:"_id" bson:"_id"`
	PmtTransactionCtx            PmtTransactionCtx          `json:"pmtTransactionCtx" bson:"pmtTransactionCtx"`
	PmtTransactionCtxID          primitive.ObjectID      		`json:"pmtTransactionCtxID" bson:"pmtTransactionCtxID"`
	PmtTransactionInitiatorCtx   PmtTransactionInitiatorCtx `json:"pmtTransactionInitiatorCtx" bson:"pmtTransactionInitiatorCtx"`
	PmtTransactionInitiatorCtxID primitive.ObjectID      		`json:"pmtTransactionInitiatorCtxID" bson:"pmtTransactionInitiatorCtxID"`
	TransactionReference      	 string                  		`json:"transactionReference" bson:"transactionReference"`
	PaymentClientReference    	 string                  		`json:"paymentClientReference" bson:"paymentClientReference"`
	PaymentClient             	 PaymentClient           		`json:"paymentClient" bson:"paymentClient"`
	PmtTransactionStatus         PmtTransactionStatus       `json:"pmtTransactionStatus" bson:"pmtTransactionStatus"`
	Amount                    	 float64                 		`json:"amount" bson:"amount"`
	Fee                       	 float64                 		`json:"fee" bson:"fee"`
	Meta                      	 map[string]interface{}  		`json:"meta" bson:"meta"`
	CreatedAt                 	 time.Time               		`json:"createdAt" bson:"createdAt"`
	UpdatedAt                 	 time.Time               		`json:"updatedAt" bson:"updatedAt"`
}

var TransactionModel *Model[Transaction]
var TransactionCollection *mongo.Collection

func initTransaction() {
	TransactionCollection = config.GetCollection(string(ModelNamesTransactions))
	TransactionModel = InitModel[Transaction](ModelNamesTransactions, TransactionCollection)
}
