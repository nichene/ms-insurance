package postgres

import (
	"context"
	"testing"

	"ms-insurance/config"
	"ms-insurance/internal/product"
	testConfig "ms-insurance/pkg/postgres"

	"github.com/stretchr/testify/assert"
)

// TestProductRepo tests all functions related to the product repository.
// The tests are made by connecting to a testcontainer.
// Should setup and cleanup test data.
func TestProductRepo(t *testing.T) {
	ctx := context.Background()
	cfg := config.LoadEnvVars(config.WithConfigPath(testConfig.ProjectRootDir))

	// >>>> Tests setup <<<<
	integrationTestDbConfig := testConfig.NewTestDb(t, ctx, cfg)

	conn := integrationTestDbConfig.InitDatabase()

	productRepository := NewProductsRepository(conn)

	// >>>> Create Product DB tests <<<<
	t.Run("Should create product find and return it", func(t *testing.T) {
		intent := &product.Product{
			Name:          "test name",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.2,
		}

		createdProduct, err := productRepository.Create(ctx, intent)
		assert.Nil(t, err)

		foundProduct, err := productRepository.FindByName(ctx, createdProduct.Name)

		assert.Nil(t, err)
		assert.NotNil(t, foundProduct.ID)
		assert.Equal(t, foundProduct.Name, intent.Name)
		assert.Equal(t, foundProduct.Category, intent.Category)
		assert.Equal(t, foundProduct.BasePrice, intent.BasePrice)
		assert.Equal(t, foundProduct.TariffedPrice, intent.TariffedPrice)
	})

	t.Run("Should create product update and return it", func(t *testing.T) {
		intent := &product.Product{
			Name:          "test name 2",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.2,
		}

		createdProduct, err := productRepository.Create(ctx, intent)
		assert.Nil(t, err)

		createdProduct.BasePrice = 500.0
		createdProduct.TariffedPrice = 600.0

		foundProduct, err := productRepository.Update(ctx, createdProduct)

		assert.Nil(t, err)
		assert.NotNil(t, foundProduct.ID)
		assert.Equal(t, foundProduct.Name, intent.Name)
		assert.Equal(t, foundProduct.Category, intent.Category)
		assert.Equal(t, foundProduct.BasePrice, createdProduct.BasePrice)
		assert.Equal(t, foundProduct.TariffedPrice, createdProduct.TariffedPrice)
	})

	// >>>> Testes teardown <<<<
	integrationTestDbConfig.ClearDatabase()
}
