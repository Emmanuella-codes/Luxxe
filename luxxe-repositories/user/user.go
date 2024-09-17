package user

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/go-kit/log"
)

type UserRespository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	QueryByID(ctx context.Context, ID string) (*entities.User, error)
	QueryByEmail(ctx context.Context, email string) (*entities.User, error)
	UpdatePassword(ctx context.Context, ID string, password string) (*entities.User, error)
	UpdateEmailYetToBeVerified(ctx context.Context, userID string, emailYetToBeVerified string) (*entities.User, error)
	VerifyUser(ctx context.Context, ID string) (*entities.User, error)
}

var UserRepo UserRespository

func InitUserRepo(logger *log.Logger) {
	UserRepo = newMgRepository(logger)
}
