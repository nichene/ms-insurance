package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "ms-insurance/docs"
)

// swagger
func SwagRoute(route *fiber.App) {
	route.Get("/swagger/*", swagger.HandlerDefault)
}
