package services

import (
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/typings"
)

var BaseAuthToken = authTokenService(&AuthTokenStruct{
	authPolicy:          headerBearerToken,
	allowExternalAccess: false,
})

func IsAnyUser(token *AccountTokenStruct) typings.FiberMiddleware {
	return accessControlService(&AccessControlStruct{
		AccountTypes: []typings.AccountType{typings.AccountTypeUser},
		AccountRoles: []entities.AccountRole{
			entities.AccountRoleUser,
			entities.AccountRoleAdmin,
		},
		Token: token,
	})
}

var IsAnyUserMiddleware = IsAnyUser(&AccountTokenStruct{})
