package repositories

import (
	"os"

	"github.com/Emmanuella-codes/Luxxe/luxxe-repositories/product"
	"github.com/Emmanuella-codes/Luxxe/luxxe-repositories/tempstore"
	"github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	"github.com/go-kit/log"
)

func InitRepositories() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	user.InitUserRepo(&logger)
	tempstore.InitTempStoreRepo(&logger)
	product.InitProductRepo(&logger)
}
