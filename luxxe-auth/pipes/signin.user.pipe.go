package pipes

import (
	"context"
	"fmt"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

func SignInUserPipe(ctx context.Context, dto *dtos.SignInUserDTO) *shared.PipeRes[entities.User] {
	email, password := dto.Email, dto.Password

	user, err := repo_user.UserRepo.QueryByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.INCORRECT_PASSWORD_EMAIL,
		}
	}

	hashedPassword := user.Password
	passwordVerified := services.ComparePasswords(hashedPassword, password)
	if !passwordVerified {
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.INCORRECT_PASSWORD_EMAIL,
		}
	}

	token, _ := services.IssueToken(&services.AccountTokenStruct{
		UserID:      user.ID.Hex(),
		AccountRole: user.AccountRole,
		Email:       user.Email,
		AccountType: typings.AccountTypeUser,
	})

	return &shared.PipeRes[entities.User]{
		Success:  true,
		Message:  messages.SUCCESS_SIGN_IN,
		Data:     user,
		HookData: user,
		Token:    token,
	}
}
