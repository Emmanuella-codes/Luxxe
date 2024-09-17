package user

import (
	"context"
	"strings"
	"time"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
)

type mgRepository struct {
	log *log.Logger
}

func newMgRepository(log *log.Logger) UserRespository {
	return &mgRepository{
		log: log,
	}
}

func (r *mgRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	firstname, lastname, email, hashedPassword :=
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password

	return entities.UserModel.Create(
		ctx,
		&bson.M{
			"firstname":   firstname,
			"lastname":    lastname,
			"email":       strings.ToLower(email),
			"password":    hashedPassword,
			"createdAt":   time.Now(),
		},
	)
}

func (r *mgRepository) QueryByID(ctx context.Context, ID string) (*entities.User, error) {
	return entities.UserModel.FindById(ctx, ID)
}

func (r *mgRepository) QueryByEmail(ctx context.Context, email string) (*entities.User, error) {
	return entities.UserModel.FindOne(ctx, &bson.M{"email": strings.ToLower(email)})
}

func (r *mgRepository) UpdatePassword(ctx context.Context, ID string, hashedPassword string) (*entities.User, error) {
	return entities.UserModel.FindByIdAndUpdate(
		ctx,
		ID,
		&bson.M{
			"$set": bson.M{
				"password":  hashedPassword,
				"updatedAt": time.Now(),
			},
		},
	)
}

func (r *mgRepository) VerifyUser(ctx context.Context, userID string) (*entities.User, error) {
	return entities.UserModel.FindByIdAndUpdate(ctx,
		userID,
		&bson.M{
			"$set": bson.M{
				"isVerified": true,
				"verifiedAt": time.Now(),
				"updatedAt":  time.Now(),
			},
		},
	)
}

func (r *mgRepository) UpdateEmailYetToBeVerified(ctx context.Context, userID string, emailYetToBeVerified string) (*entities.User, error) {
	return entities.UserModel.FindByIdAndUpdate(ctx,
		userID,
		&bson.M{
			"$set": bson.M{
				"emailYetToBeVerified": emailYetToBeVerified,
				"updatedAt":            time.Now(),
			},
		},
	)
}
