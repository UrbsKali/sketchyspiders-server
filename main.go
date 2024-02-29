package main

import (
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
	server()
}
