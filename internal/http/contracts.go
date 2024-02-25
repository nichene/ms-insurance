package http

import (
	"ms-insurance/internal/health"
	"ms-insurance/internal/product"
)

type Service struct {
	Health         health.Checks
	ProductService product.Service
}

func NewService(health health.Checks, productService product.Service) *Service {
	return &Service{
		Health:         health,
		ProductService: productService,
	}
}
