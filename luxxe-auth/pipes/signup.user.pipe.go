package pipes

import (
	"context"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared "github.com/Emmanuella-codes/Luxxe/luxxe-shared"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

func SignUpUserPipe(ctx context.Context, dto *dtos.SignUpUserDTO) *shared.PipeRes[entities.User] {
	firstname, lastname, email, password, adminKey := dto.Firstname, dto.Lastname, dto.Email, dto.Password, dto.AdminKey

	userExists, _ := repo_user.UserRepo.QueryByEmail(ctx, email)
	if userExists != nil {
		return &shared.PipeRes[entities.User]{
			Success: false,
			Message: messages.EXISTING_ACCOUNT_REGISTERED_EMAIL,
		}
	}

	hashedPassword := services.GeneratePasswordHash(password)

	var accountRole entities.AccountRole
	if adminKey != "" && adminKey == config.EnvConfig.ADMIN_KEY {
		accountRole = entities.AccountRoleAdmin
	} else {
		accountRole = entities.AccountRoleUser
	}

	user, _ := repo_user.UserRepo.Create(
		ctx,
		&entities.User{
			Firstname: 	 firstname,
			Lastname: 	 lastname,
			Email:    	 email,
			Password:    hashedPassword,
			AccountRole: accountRole,
		})

	token, _ := services.IssueToken(&services.AccountTokenStruct{
		UserID:      user.ID.Hex(),
		AccountRole: accountRole,
		Email:       user.Email,
		AccountType: typings.AccountTypeUser,
	})

	return &shared.PipeRes[entities.User]{
		Success:  true,
		Message:  messages.CREATED_NEW_ACCOUNT,
		Data:     user,
		HookData: user,
		Token:    token,
	}
}
