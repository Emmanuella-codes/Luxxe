package transactions

import (
	"context"
	"time"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/misc"
	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) TransactionRepository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) Create(ctx context.Context, transaction *entities.Transaction) (*entities.Transaction, error) {
	pmtTransactionCtx            := transaction.PmtTransactionCtx
	pmtTransactionCtxID          := transaction.PmtTransactionCtxID 
	pmtTransactionInitiatorCtx 	 := transaction.PmtTransactionInitiatorCtx
	pmtTransactionInitiatorCtxID := transaction.PmtTransactionInitiatorCtxID
	transactionReference         := transaction.TransactionReference
	paymentClient                := transaction.PaymentClient
	amount                       := transaction.Amount

	return entities.TransactionModel.Create(
		ctx, 
		&bson.M{
			"pmtTransactionCtx":   					pmtTransactionCtx,
			"pmtTransactionCtxID": 					pmtTransactionCtxID,
			"pmtTransactionInitiatorCtx": 	pmtTransactionInitiatorCtx,
			"pmtTransactionInitiatorCtxID": pmtTransactionInitiatorCtxID,
			"transactionReference":         transactionReference,
			"transactionStatus":         		entities.PmtTxStatusPending,
			"paymentClient":                paymentClient,
			"amount":                       amount,
			"createdAt":                 		time.Now(),
		},
	)
}

func (r *mgRepository) QueryPendingTransactionByTransactionReference(ctx context.Context, transactionReference string) (*entities.Transaction, error) {
	return entities.TransactionModel.FindOne(
		ctx,
		&bson.M{
			"transactionReference":         transactionReference,
			"transactionStatus":         		entities.PmtTxStatusPending,
		},
	)
}

func (r *mgRepository) QueryAllTransactionsByTransactionInitiatorCtxID(
	ctx context.Context,
	pmtTransactionCtx entities.PmtTransactionCtx,
	pmtTransactionInitiatorCtxIDStr string,
	page int,
) (*[]entities.Transaction, int64, error) {
	pmtTransactionInitiatorCtxIDObj, _ := primitive.ObjectIDFromHex(pmtTransactionInitiatorCtxIDStr)

	skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := bson.M{
		"pmtTransactionCtx":   					pmtTransactionCtx,
		"pmtTransactionInitiatorCtxID": pmtTransactionInitiatorCtxIDObj,
	}

	transactionCount, _ := entities.TransactionCollection.CountDocuments(ctx, &filter)

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	cursor, err := entities.TransactionCollection.Find(ctx, &filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var transactions []entities.Transaction = []entities.Transaction{}
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, 0, err
	}
	return &transactions, transactionCount, nil
}

func (r *mgRepository) QueryTransactionsByOrderID(
	ctx context.Context,
	orderID string,
	page int,
	transactionStatus entities.PmtTransactionStatus,
) (*[]entities.Transaction, int64, error) {
	skip, limit := misc.Pagination(misc.PaginationStruct{Page: page})

	filter := bson.M{
		"meta.orderID": orderID,
		"pmtTransactionCtx": entities.PmtTxCtxOrder,
	}
	if transactionStatus != "" {
		filter["transactionStatus"] = transactionStatus
	}

	transactionCount, _ := entities.TransactionCollection.CountDocuments(ctx, filter)
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))

	cursor, err := entities.TransactionCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// var transactions []entities.Transaction
	transactions := make([]entities.Transaction, 0)
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, 0, err
	}

	return &transactions, transactionCount, nil
}

func (r *mgRepository) QueryByTransactionReference(ctx context.Context, transactionReference string) (*entities.Transaction, error) {
	return entities.TransactionModel.FindOne(
		ctx,
		&bson.M{
			"transactionReference": transactionReference,
		},
	)
}

func (r *mgRepository) UpdateTransactionStatus(ctx context.Context, txID string, txStatus entities.PmtTransactionStatus) (*entities.Transaction, error) {
	return entities.TransactionModel.FindByIdAndUpdate(
		ctx,
		txID,
		&bson.M{
			"$set": &bson.M{
				"pmtTransactionStatus": txStatus,
				"updatedAt":         		time.Now(),
			},
		},
	)
}

func (r *mgRepository) UpdatePaymentClientReference(ctx context.Context, txID string, paymentClientReference string) (*entities.Transaction, error) {
	return entities.TransactionModel.FindByIdAndUpdate(
		ctx,
		txID,
		&bson.M{
			"$set": &bson.M{
				"paymentClientReference": paymentClientReference,
				"updatedAt":              time.Now(),
			},
		},
	)
}
