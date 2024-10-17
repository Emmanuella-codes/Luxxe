package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
)

func ResetUserPasswordByEmailPipe(ctx context.Context, dto *dtos.ResetUserPasswordByEmailDTO) *shared.PipeRes[entities.User] {
	email, otp, password := dto.Email, dto.Otp, dto.Password

	userExists, err := repo_user.UserRepo.QueryByEmail(ctx, email)
	if err != nil {
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.NOT_FOUND_USER,
		}
	}

	getOtp := services.VerifyOtp(ctx, &services.VerifyOtpStruct{Email: email, Otp: otp})
	if !getOtp {
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.INCORRECT_EXPIRED_OTP,
		}
	}

	hashedPassword := services.GeneratePasswordHash(password)
	repo_user.UserRepo.UpdatePassword(ctx, userExists.ID.Hex(), hashedPassword)

	return &shared.PipeRes[entities.User]{
		Success:  true,
		Message:  messages.SUCCESS_CHANGE_PASSWORD,
		Data:     userExists,
		HookData: userExists,
	}
}
