package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"ms-insurance/config"
	"ms-insurance/internal/health"
	"ms-insurance/internal/product"
	productDB "ms-insurance/internal/product/postgres"
	testConfig "ms-insurance/pkg/postgres"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// >>>> E2E tests <<<<

// TestE2E should test all endpoints.
// The tests are made by starting a new server and connecting to a testcontainer.
// Should setup and cleanup test data.
func TestE2E(t *testing.T) {

	// >>>> Tests setup <<<<
	testingTimeout := 30000 // 30s
	ctx := context.Background()
	cfg := config.LoadEnvVars(config.WithConfigPath(testConfig.ProjectRootDir))

	integrationTestDbConfig := testConfig.NewTestDb(t, ctx, cfg)

	conn := integrationTestDbConfig.InitDatabase()

	productRepository := productDB.NewProductsRepository(conn)

	httpService := NewService(
		health.NewHealthCheckService(),
		*product.NewService(productRepository),
	)

	srv, err := NewServer(cfg.Port, WithService(httpService))
	assert.Nil(t, err)

	//  >>>> Create product endpoint tests <<<<

	t.Run("Create product endpoint success", func(t *testing.T) {
		endpoint := "/product"
		payloadStr := "{\"name\": \"Teste\", \"category\": \"VIDA\", \"base_price\": 100.00}"
		payload := strings.NewReader(payloadStr)
		request := httptest.NewRequest(http.MethodPost, endpoint, payload)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		var createRespose product.Product
		body, err := io.ReadAll(response.Body)
		json.Unmarshal(body, &createRespose)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, createRespose.Name, "Teste")
		assert.Equal(t, createRespose.Category, "VIDA")
		assert.Equal(t, createRespose.BasePrice, 100.0)
		assert.Equal(t, createRespose.TariffedPrice, 103.2)
	})

	t.Run("Create product endpoint bad request", func(t *testing.T) {
		endpoint := "/product"
		payloadStr := "{\"name\": \"Teste\"}"
		payload := strings.NewReader(payloadStr)
		request := httptest.NewRequest(http.MethodPost, endpoint, payload)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	//  >>>> FindByName product endpoint tests <<<<

	t.Run("FindByName product endpoint success", func(t *testing.T) {
		expectedProduct := product.Product{
			Name:          "Teste",
			Category:      "VIDA",
			BasePrice:     100.0,
			TariffedPrice: 103.20,
		}

		endpoint := "/product/?name=Teste"
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		var findRespose product.Product
		body, err := io.ReadAll(response.Body)
		json.Unmarshal(body, &findRespose)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, findRespose.Name, expectedProduct.Name)
		assert.Equal(t, findRespose.Category, expectedProduct.Category)
		assert.Equal(t, findRespose.BasePrice, expectedProduct.BasePrice)
		assert.Equal(t, findRespose.TariffedPrice, expectedProduct.TariffedPrice)
	})

	t.Run("FindByName product endpoint query param error", func(t *testing.T) {
		endpoint := "/product"
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	//  >>>> Update product endpoint tests <<<<

	t.Run("Update product endpoint success", func(t *testing.T) {
		endpoint := "/product"
		payloadStr := "{\"name\": \"Teste2\", \"category\": \"AUTO\", \"base_price\": 100.00}"
		payload := strings.NewReader(payloadStr)
		request := httptest.NewRequest(http.MethodPost, endpoint, payload)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		var findRespose product.Product
		body, err := io.ReadAll(response.Body)
		json.Unmarshal(body, &findRespose)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		endpoint = "/product"
		payloadStr = fmt.Sprintf("{\"id\": %d, \"name\": \"Teste3\", \"category\": \"AUTO\", \"base_price\": 100.00}", findRespose.ID)
		payload = strings.NewReader(payloadStr)
		request = httptest.NewRequest(http.MethodPut, endpoint, payload)
		request.Header.Set("Content-Type", "application/json")

		response, err = srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		var updateRespose product.Product
		body, err = io.ReadAll(response.Body)
		json.Unmarshal(body, &updateRespose)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, updateRespose.Name, "Teste3")
	})

	t.Run("Update product endpoint validate error", func(t *testing.T) {
		endpoint := "/product"
		payloadStr := "{\"name\": \"Teste3\", }"
		payload := strings.NewReader(payloadStr)
		request := httptest.NewRequest(http.MethodPut, endpoint, payload)
		request.Header.Set("Content-Type", "application/json")

		response, err := srv.httpServer.Test(request, testingTimeout)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})

	// >>>> Testes teardown <<<<
	integrationTestDbConfig.ClearDatabase()
}
