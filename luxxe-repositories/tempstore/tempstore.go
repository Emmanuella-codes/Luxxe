package tempstore

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/go-kit/log"
)

type TempStoreRespository interface {
	Create(ctx context.Context, tempstore *entities.TempStore) (*entities.TempStore, error)
	QueryByKey(ctx context.Context, key string) (*entities.TempStore, error)
	UpdateKeyAndValue(ctx context.Context, key string, value string, expirationTime int, beginTime int64) (*entities.TempStore, error)
}

var TempStoreRepo TempStoreRespository

func InitTempStoreRepo(logger *log.Logger) {
	TempStoreRepo = newMgRepository(logger)
}
