package services

import (
	"fmt"

	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	"github.com/Emmanuella-codes/Luxxe/typings"
	"github.com/gofiber/fiber/v2"
)

type AccessControlStruct struct {
	AccountTypes []typings.AccountType
	AccountRoles []entities.AccountRole
	Token        *AccountTokenStruct
}

func accessControlService(acs *AccessControlStruct) typings.FiberMiddleware {
	return func(ctx *fiber.Ctx) error {
		AccountTypes, accountRoles, token := acs.AccountTypes, acs.AccountRoles, acs.Token
		var AccountToken *AccountTokenStruct
		if token.UserID != "" {
			AccountToken = token
		} else {
			AccountToken = ctx.Locals("token").(*AccountTokenStruct)
		}

		var accountTypeAccessGranted bool
		for _, actType := range AccountTypes {
			if actType == AccountToken.AccountType {
				accountTypeAccessGranted = true
				break
			}
		}

		var accountRoleAccessGranted bool
		for _, actRole := range accountRoles {
			if actRole == AccountToken.AccountRole {
				accountRoleAccessGranted = true
				break
			}
		}

		if !accountTypeAccessGranted || !accountRoleAccessGranted {
			errorMsg := ""
			if !accountTypeAccessGranted {
				errorMsg += " account type error, "
			}
			if !accountTypeAccessGranted {
				errorMsg += " account role error, "
			}

			fmt.Println(accountTypeAccessGranted, accountRoleAccessGranted)

			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"statusCode": fiber.StatusUnauthorized,
					"message":    "You are not Authorized to perform this operation! ::: " + errorMsg,
				},
			)
		}

		// TODO: some other conditions
		return ctx.Next()
	}
}

