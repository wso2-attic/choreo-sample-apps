package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"example.choreo.dev/internal/config"
)

// @Summary      Get pet info by id
// @Produce      json
// @Router /healthz [post]
// @Success 200 {object} map[string]any "ok"
func handleHealthCheckRequest(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message":   "choreo-example-app is healthy",
		"env":       config.GetConfig().Env,
		"timestamp": time.Now(),
	})
}

func RegisterHealthRoutes(r fiber.Router) {
	r.Get("/healthz", handleHealthCheckRequest)
}
