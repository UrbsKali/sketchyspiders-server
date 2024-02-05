package routes

import (
	"urbskali/ssserver/api"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {

	app.Post("/login", api.Login)
	app.Get("/current_version", api.GetCurrentVersion)
	app.Get("/notify", api.GetNotify)
}
