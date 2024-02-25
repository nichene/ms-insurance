package routes

import (
	"ms-insurance/internal/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// ProductRoutes contains product endpoints.
func ProductRoutes(route *fiber.App, handler handlers.ProductHandler) {
	route.Post("/product", handler.Create)
	route.Get("/product", handler.FindByName)
	route.Patch("/product", handler.Update)
}
