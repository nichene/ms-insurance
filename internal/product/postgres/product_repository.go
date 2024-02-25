package postgres

import (
	"context"
	"ms-insurance/internal/product"
	"time"

	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{db}
}

func (r *ProductsRepository) FindByName(ctx context.Context, name string) (*product.Product, error) {
	foundProduct := &product.Product{}

	result := r.db.Model(&product.Product{}).Where("name = ?", name).First(foundProduct)

	return foundProduct, result.Error
}

func (r *ProductsRepository) Create(ctx context.Context, intent *product.Product) (*product.Product, error) {
	intent.CreatedAt = time.Now().UTC()
	intent.UpdatedAt = time.Now().UTC()

	err := r.db.Create(&intent).Error
	if err != nil {
		return &product.Product{}, err
	}

	foundProduct := &product.Product{}
	err = r.db.Model(&product.Product{}).Where("name = ?", intent.Name).First(foundProduct).Error

	return foundProduct, err
}

func (r *ProductsRepository) Update(ctx context.Context, intent *product.Product) (*product.Product, error) {
	intent.UpdatedAt = time.Now().UTC()

	err := r.db.Model(&product.Product{}).Where("id = ?", intent.ID).Updates(&intent).Error
	if err != nil {
		return &product.Product{}, err
	}

	foundProduct := &product.Product{}
	err = r.db.Model(&product.Product{}).Where("name = ?", intent.Name).First(foundProduct).Error

	return foundProduct, err
}
