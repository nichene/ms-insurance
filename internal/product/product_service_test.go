package product

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	productsRepoMock := NewMockProductsRepository(ctrl)

	service := NewService(productsRepoMock)

	t.Run("Should create a product", func(t *testing.T) {
		product := &Product{
			Name:          "Seguro de Vida Individual",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.2,
		}

		productsRepoMock.EXPECT().Create(ctx, product).Return(product, nil)

		createdProduct, err := service.Create(ctx, product)

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, product, createdProduct)
	})

	t.Run("Should not create a product when Categoory is incorrect", func(t *testing.T) {
		product := &Product{
			Name:      "Seguro de Vida Individual",
			Category:  "VIDAs",
			BasePrice: 100.0,
		}

		createdProduct, err := service.Create(ctx, product)

		assert.ErrorContains(t, err, "Unable to create product with category: VIDAs")
		assert.Equal(t, 0, createdProduct.ID)
	})
}

func TestFindProductByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	productsRepoMock := NewMockProductsRepository(ctrl)

	service := NewService(productsRepoMock)

	t.Run("Should find a product", func(t *testing.T) {
		product := &Product{
			Name:          "Seguro de Vida Individual",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.2,
		}

		productsRepoMock.EXPECT().FindByName(ctx, product.Name).Return(product, nil)

		createdProduct, err := service.FindByName(ctx, product.Name)

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, product, createdProduct)
	})
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	productsRepoMock := NewMockProductsRepository(ctrl)

	service := NewService(productsRepoMock)

	t.Run("Should update a product", func(t *testing.T) {
		product := &Product{
			Name:          "Seguro de Vida Individual",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.2,
		}

		productsRepoMock.EXPECT().Update(ctx, product).Return(product, nil)

		createdProduct, err := service.Update(ctx, product)

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, product, createdProduct)
	})

	t.Run("Should not update a product when Categoory is incorrect", func(t *testing.T) {
		product := &Product{
			Name:      "Seguro de Vida Individual",
			Category:  "VIDAs",
			BasePrice: 100.0,
		}

		createdProduct, err := service.Update(ctx, product)

		assert.ErrorContains(t, err, "Unable to update product with category: VIDAs")
		assert.Equal(t, 0, createdProduct.ID)
	})
}
