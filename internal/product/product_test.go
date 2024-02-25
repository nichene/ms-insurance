package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsIdentifiableCategory(t *testing.T) {
	t.Run("Should return true for identifiable category VIDA", func(t *testing.T) {
		product := &Product{Category: "VIDA"}

		isCategory := product.containsIdentifiableCategory()

		assert.True(t, isCategory)
	})

	t.Run("Should return true for identifiable category VIAGEM", func(t *testing.T) {
		product := &Product{Category: "VIAGEM"}

		isCategory := product.containsIdentifiableCategory()

		assert.True(t, isCategory)
	})

	t.Run("Should return true for identifiable category AUTO", func(t *testing.T) {
		product := &Product{Category: "AUTO"}

		isCategory := product.containsIdentifiableCategory()

		assert.True(t, isCategory)
	})

	t.Run("Should return true for identifiable category RESIDENCIAL", func(t *testing.T) {
		product := &Product{Category: "RESIDENCIAL"}

		isCategory := product.containsIdentifiableCategory()

		assert.True(t, isCategory)
	})

	t.Run("Should return true for identifiable category PATRIMONIAL", func(t *testing.T) {
		product := &Product{Category: "PATRIMONIAL"}

		isCategory := product.containsIdentifiableCategory()

		assert.True(t, isCategory)
	})

	t.Run("Should return false for non identifiable categories", func(t *testing.T) {
		product := &Product{
			Category: "VIDAs",
		}

		isCategory := product.containsIdentifiableCategory()

		assert.False(t, isCategory)
	})
}

func TestCalculateInsuranceProductTariff(t *testing.T) {
	t.Run("Should calculate for category VIDA", func(t *testing.T) {
		product := &Product{
			Category:  "VIDA",
			BasePrice: 100.0,
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 103.20, product.TariffedPrice)
	})

	t.Run("Should calculate for category VIAGEM", func(t *testing.T) {
		product := &Product{
			Category:  "VIAGEM",
			BasePrice: 100,
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 107.0, product.TariffedPrice)
	})

	t.Run("Should calculate for category AUTO", func(t *testing.T) {
		product := &Product{
			Category:  "AUTO",
			BasePrice: 50.0,
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 55.25, product.TariffedPrice)
	})

	t.Run("Should calculate for category RESIDENCIAL", func(t *testing.T) {
		product := &Product{
			Category:  "RESIDENCIAL",
			BasePrice: 100.0,
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 107.0, product.TariffedPrice)
	})

	t.Run("Should calculate for category PATRIMONIAL", func(t *testing.T) {
		product := &Product{
			Category:  "PATRIMONIAL",
			BasePrice: 100.0,
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 108.0, product.TariffedPrice)
	})

	t.Run("Should not calculate for non identifiable categories", func(t *testing.T) {
		product := &Product{
			Category: "VIDAs",
		}

		product.calculateInsuranceProductTariff()

		assert.Equal(t, 0.0, product.TariffedPrice)
	})
}
