package pipes

import (
	"context"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/luxxe-profile/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-profile/messages"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func GetUserProfilePipe(ctx context.Context, dto *dtos.GetUserProfileDTO) *shared.PipeRes[entities.User] {
	userID := dto.UserID

	user, err :=repo_user.UserRepo.QueryByID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.FAILURE_USER_NOT_FETCHED,
		}
	}

	return &shared.PipeRes[entities.User]{
		Success: true,
		Message: messages.SUCCESS_USER_FETCHED,
		Data: user,
	}
}
