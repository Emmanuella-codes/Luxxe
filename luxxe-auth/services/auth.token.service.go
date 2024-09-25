package services

import (
	"github.com/Emmanuella-codes/Luxxe/typings"
	"github.com/gofiber/fiber/v2"
)

type AuthTokenStruct struct {
	authPolicy          typings.TokenGeneratorFunc
	allowExternalAccess bool
}

func authTokenService(ats *AuthTokenStruct) typings.FiberMiddleware {
	authPolicy, allowExternalAccess := ats.authPolicy, ats.allowExternalAccess
	return func(ctx *fiber.Ctx) error {
		tokenString := authPolicy(&typings.TokenGen{Ctx: ctx})

		if allowExternalAccess && tokenString == "" {
			return ctx.Next()
		} else if tokenString != "" {
			verifiedToken, err := VerifyToken(tokenString)
			if err != nil {
				if allowExternalAccess {
					return ctx.Next()
				}
				return ctx.Status(fiber.StatusUnauthorized).JSON(
					fiber.Map{
						"statusCode": fiber.StatusUnauthorized,
						"message":    "Invalid Token",
					},
				)
			}

			ctx.Locals("token", verifiedToken)
			return ctx.Next()
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"message":    "No Authorization found",
			},
		)
	}
}
