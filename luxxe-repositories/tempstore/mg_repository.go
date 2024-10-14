package tempstore

import (
	"context"
	"time"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) TempStoreRespository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) Create(ctx context.Context, tempstore *entities.TempStore) (*entities.TempStore, error) {
	key, value, expirationTime, beginTime := tempstore.Key, tempstore.Value, tempstore.ExpirationTime, tempstore.BeginTime

	return entities.TempStoreModel.Create(
		ctx,
		&bson.M{
			"key":            key,
			"value":          value,
			"expirationTime": expirationTime,
			"beginTime":      beginTime,
			"createdAt":      time.Now(),
		},
	)
}

func (r *mgRepository) QueryByKey(ctx context.Context, key string) (*entities.TempStore, error) {
	return entities.TempStoreModel.FindOne(ctx, &bson.M{"key": key})
}

func (r *mgRepository) UpdateKeyAndValue(ctx context.Context, key string, value string, expirationTime int, beginTime int64) (*entities.TempStore, error) {
	return entities.TempStoreModel.FindOneAndUpdate(
		ctx,
		&bson.M{"key": key},
		&bson.M{
			"$set": &bson.M{
				"key":            key,
				"value":          value,
				"expirationTime": expirationTime,
				"beginTime":      beginTime,
				"updatedAt":      time.Now(),
			},
		},
	)
}
