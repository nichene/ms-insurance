package product

import (
	"context"
)

func (s *Service) Create(ctx context.Context, intent Product) (product *Product, err error) {
	intent.calculateInsuranceProductTariff()

	return s.productsRepository.Create(ctx, &intent)
}

func (s *Service) FindByName(ctx context.Context, name string) (product *Product, err error) {
	return s.productsRepository.FindByName(ctx, name)
}

func (s *Service) Update(ctx context.Context, name string) (product *Product, err error) {
	//TODO
	return product, err
}
