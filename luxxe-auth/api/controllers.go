package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/dtos"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/messages"
	"github.com/Emmanuella-codes/Luxxe/luxxe-auth/services"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repo_user "github.com/Emmanuella-codes/Luxxe/luxxe-repositories/user"
	shared_api "github.com/Emmanuella-codes/Luxxe/luxxe-shared/api"
	"github.com/Emmanuella-codes/Luxxe/typings"
)