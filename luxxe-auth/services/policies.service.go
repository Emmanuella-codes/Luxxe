package services

import (
	"strings"

	"github.com/Emmanuella-codes/Luxxe/typings"
)

func headerBearerToken(tkg *typings.TokenGen) string {
	ctx := tkg.Ctx
	authHeader := ctx.Get("Authorization")
	return strings.TrimPrefix(authHeader, "Bearer ")
}