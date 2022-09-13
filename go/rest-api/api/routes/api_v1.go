package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Initialize(app *fiber.App) {
	initControllers()

	RegisterHealthRoutes(app)
	apiVersion := app.Group("/api/v1")
	registerPetRoutes(apiVersion)
	registerCategoryRoutes(apiVersion)
}
