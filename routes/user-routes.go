package routes

import (
	"urbskali/ssserver/api"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	route := app.Group("/api")

	route.Post("/delete/", api.Delete)
	route.Post("/upload/", api.Upload)
	route.Post("/create_dir/", api.CreateDir)
}
