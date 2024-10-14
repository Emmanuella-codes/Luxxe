package luxxestorage

import (
	"context"
	"errors"
	"time"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_tempstore "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/tempstore"
)

type SetStruct struct {
	Key            string
	Value          string
	ExpirationTime int
}

func Set(ctx context.Context, setStrPtr *SetStruct) {
	key, value, expirationTime := setStrPtr.Key, setStrPtr.Value, setStrPtr.ExpirationTime

	_, err := repo_tempstore.TempStoreRepo.QueryByKey(ctx, key)
	currentTime := time.Now()
	beginTime := currentTime.UnixNano() / int64(time.Millisecond)
	if err != nil {
		repo_tempstore.TempStoreRepo.Create(ctx, &entities.TempStore{
			Key:            key,
			Value:          value,
			ExpirationTime: expirationTime,
			BeginTime:      beginTime,
		})
	} else {
		repo_tempstore.TempStoreRepo.UpdateKeyAndValue(ctx, key, value, expirationTime, beginTime)
	}
}

func Get(ctx context.Context, key string) (string, error) {
	tstore, err := repo_tempstore.TempStoreRepo.QueryByKey(ctx, key)
	if err != nil {
		return "", err
	}
	value, expirationTime, beginTime := tstore.Value, tstore.ExpirationTime, tstore.BeginTime
	if expirationTime != 0 {
		then := beginTime
		currentTime := time.Now()
		now := currentTime.UnixNano() / int64(time.Millisecond)
		timeElapsed := now > then+(int64(expirationTime)*1000)
		if timeElapsed {
			return "", errors.New("elapsed time")
		}
	}
	return value, nil
}
