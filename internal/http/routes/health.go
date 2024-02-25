package routes

import (
	"ms-insurance/internal/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// HealthRoute contains health check endpoints.
func HealthRoute(route *fiber.App, handler handlers.HealthHandler) {
	route.Get("/health", handler.Check)
}
