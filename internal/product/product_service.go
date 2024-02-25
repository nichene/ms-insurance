package product

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func (s *Service) Create(ctx context.Context, intent *Product) (product *Product, err error) {
	if !intent.containsIdentifiableCategory() {
		msg := fmt.Sprintf("Unable to create product with category: %s", intent.Category)
		log.Default().Print(msg)
		return &Product{}, errors.New(msg)
	}

	intent.calculateInsuranceProductTariff()

	return s.productsRepository.Create(ctx, intent)
}

func (s *Service) FindByName(ctx context.Context, name string) (product *Product, err error) {
	return s.productsRepository.FindByName(ctx, name)
}

func (s *Service) Update(ctx context.Context, intent *Product) (product *Product, err error) {
	if !intent.containsIdentifiableCategory() {
		msg := fmt.Sprintf("Unable to update product with category: %s", intent.Category)
		log.Default().Print(msg)
		return &Product{}, errors.New(msg)
	}

	intent.calculateInsuranceProductTariff()

	return s.productsRepository.Update(ctx, intent)
}
