package pipes

import (
	"context"
	"fmt"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/luxxe-shared/mail"
)

func SendOTPPipe(ctx context.Context, dto *dtos.SendOTPDTO) *shared.PipeRes[string] {
	email := dto.Email

	userExists, err := repo_user.UserRepo.QueryByEmail(ctx, email)
	if err != nil {
		return &shared.PipeRes[string]{
			Success: false,
			Message: messages.NOT_REGISTERED_EMAIL,
		}
	}

	ioRes := services.IssueOtp(ctx, &services.IssueOTPStruct{Email: userExists.Email})
	mail.SendOtpMail(email, ioRes.EmailOTP, fmt.Sprintf("%d minutes left", ioRes.TimeLeft))

	return &shared.PipeRes[string]{
		Success: true,
		Message: messages.SENT_OTP_EMAIL,
		Data:    &ioRes.EmailOTP,
	}
}
