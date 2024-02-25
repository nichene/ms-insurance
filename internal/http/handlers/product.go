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
// @Summary      Create product
// @Description  Create product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        request body product.Product true "Body"
// @Success      200 {object} product.Product
// @Failure      400
// @Failure      500
// @Router       /product [post]
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

	product, err := h.productService.Create(ctx.Context(), params)
	if err != nil {
		log.Default().Print("API - Error creating person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(product)
}

// Create godoc
// @Summary      FindByName product
// @Description  FindByName product
// @Tags         Product
// @Produce      json
// @Param        name   query      string  true  "name"
// @Success      200 {object} product.Product
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /product [get]
func (h *productHandler) FindByName(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	if name == "" {
		log.Default().Print("API - Unable to get query param name")
		return ctx.Status(http.StatusBadRequest).JSON("Invalid query param : name")
	}

	product, err := h.productService.FindByName(ctx.Context(), name)
	if err != nil {
		log.Default().Print("API - Error finding product")
		return ctx.Status(http.StatusNotFound).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(product)
}

// Create godoc
// @Summary      UpdateByID product
// @Description  UpdateBYID product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        request body product.Product true "Body"
// @Success      200 {object} product.Product
// @Failure      404
// @Failure      500
// @Router       /product [put]
func (h *productHandler) Update(ctx *fiber.Ctx) error {
	params := &product.Product{}
	err := ctx.BodyParser(params)
	if err != nil {
		log.Default().Print("API - Unable to parse request body", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON("Unable to parse request body")
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		log.Default().Print("API - Invalid update product params")
		return ctx.Status(http.StatusBadRequest).JSON("Invalid update product params")
	}

	product, err := h.productService.Update(ctx.Context(), params)
	if err != nil {
		log.Default().Print("API - Error creating person")
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(product)
}
