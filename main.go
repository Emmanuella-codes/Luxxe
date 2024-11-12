package main

import (
	"fmt"
	// "os"

	config "github.com/Emmanuella-codes/Luxxe/luxxe-config"
	entities "github.com/Emmanuella-codes/Luxxe/luxxe-entities"
	repositories "github.com/Emmanuella-codes/Luxxe/luxxe-repositories"
	// "github.com/Emmanuella-codes/Luxxe/luxxe-shared/bootstrap"
	web_api "github.com/Emmanuella-codes/Luxxe/luxxe-web-api"
)

func main() {
	config.InitServer()
	entities.InitModels()
	repositories.InitRepositories()

	defer config.TerminateServer()

	// bootstrap section ---> uncomment to bootstrap new folder
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run bootstrap.go <folder-name>")
	// 	return
	// }
	// folderName := os.Args[1]
	// bootstrap.BootstrapProject(folderName)

	app := web_api.GenerateApp()

	if err := app.Listen(":" + config.EnvConfig.PORT); err != nil {
		fmt.Println("Failed to start web server:", err)
	}
}
