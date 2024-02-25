package handlers

import (
	"log"
	"ms-insurance/internal/product"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	Create(ctx *fiber.Ctx) error
	FindByName(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

// Create godoc
// @Summary      Create integrator
// @Description  Create integrator
// @Tags         Integrator
// @Accept       json
// @Produce      json
// @Param        request body product.Product true "Body"
// @Success      201 {object} product.Product
// @Failure      400
// @Failure      422
// @Failure      500
// @Router       /api/management/product [post]
func (h *productHandler) Create(ctx *fiber.Ctx) error {
	params := &product.Product{}
	err := ctx.BodyParser(params)
	if err != nil {
		log.Default().Print("API - Unable to parse request body", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON("Unable to parse request body")
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		log.Default().Print("API - Invalid create product params")
		return ctx.Status(http.StatusBadRequest).JSON("Invalid create product params")
	}

	product, err := h.productService.Create(ctx.Context(), *params)
	if err != nil {
		log.Default().Print("API - Error creating person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(product)
}

func (h *productHandler) FindByName(ctx *fiber.Ctx) error {
	var name string
	err := ctx.QueryParser(&name)
	if err != nil {
		log.Default().Print("API - Unable to parse query param", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON("Invalid query params")
	}

	product, err := h.productService.FindByName(ctx.Context(), name)
	if err != nil {
		log.Default().Print("API - Error finding product")
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(product)
}

func (h *productHandler) Update(ctx *fiber.Ctx) error {
	// TODO
	return ctx.Status(http.StatusOK).JSON("product")
}
