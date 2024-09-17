package main

import (
	"fmt"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repositories "github.com/Emmanuella-codes/Luxxe/luxxe-repositories"
	web_api "github.com/Emmanuella-codes/Luxxe/luxxe-web-api"
)

func main() {
	config.InitServer()
	entities.InitModels()
	repositories.InitRepositories()

	defer config.TerminateServer()

	app := web_api.GenerateApp()

	if err := app.Listen(":" + config.EnvConfig.PORT); err != nil {
		fmt.Println("Failed to start web server:", err)
	}
}
