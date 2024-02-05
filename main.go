package main

import (
	"os"
	"urbskali/ssserver/routes"
	"urbskali/ssserver/utils"

	"github.com/gofiber/fiber/v2"
)

func server() {
	utils.StartUp()

	app := fiber.New()

	routes.StaticRoutes(app)

	routes.PublicRoutes(app)
	routes.AdminRoutes(app)

	utils.Start(app)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "setup" {
		utils.Setup()
	} else if len(os.Args) > 1 && os.Args[1] == "build-ui" {
		utils.BuildUI()
	} else {
		server()
	}
}
