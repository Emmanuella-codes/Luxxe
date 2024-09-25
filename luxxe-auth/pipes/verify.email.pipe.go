package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func VerifyEmailPipe(ctx context.Context, dto *dtos.VerifyEmailDTO) *shared.PipeRes[string] {
	userID := dto.UserID

	existing_user, err := repo_user.UserRepo.QueryByID(ctx, userID)
	if err != nil {
		return &shared.PipeRes[string]{
			Success: false,
			Message: messages.NOT_FOUND_USER,
		}
	}

	// someone might have claimed the email just before the verification
	other_user, _ := repo_user.UserRepo.QueryByEmail(ctx, existing_user.EmailYetToBeVerified)
	if other_user != nil {
		if other_user.ID.Hex() != userID {
			return &shared.PipeRes[string]{
				Success: false,
				Message: messages.CANNOT_VERIFY_ANOTHER_USERS_EMAIL,
			}
		}
	}

	repo_user.UserRepo.VerifyUser(ctx, userID)
	return &shared.PipeRes[string]{
		Success: true,
		Message: messages.SENT_OTP_EMAIL,
	}
}
