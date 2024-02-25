//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -package=product -source=contracts.go -destination=mocks_contracts.go .
package product

import (
	"context"
)

type Service struct {
	productsRepository ProductsRepository
}

func NewService(
	productsRepository ProductsRepository,
) *Service {
	return &Service{
		productsRepository: productsRepository,
	}
}

type ProductsRepository interface {
	Create(ctx context.Context, produtc *Product) (*Product, error)
	FindByName(ctx context.Context, name string) (*Product, error)
	Update(ctx context.Context, produtc Product) (*Product, error)
}
