package handlers

import (
	"ms-insurance/internal/health"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthHandler interface {
	Check(c *fiber.Ctx) error
}

type healthHandler struct {
	healthService health.Checks
}

func NewHealthHandler(healthService health.Checks) HealthHandler {
	return &healthHandler{
		healthService: healthService,
	}
}

func (h *healthHandler) Check(c *fiber.Ctx) error {
	err := h.healthService.Ping(c.Context())
	if err != nil {
		_ = c.SendStatus(http.StatusInternalServerError)
		return c.JSON(&fiber.Map{"error: ": err.Error()})
	}

	return c.JSON(&fiber.Map{"message: ": "successful ping"})
}
