package tests

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/agustinrabini/api-prices-project/cmd/api/handlers"
	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/agustinrabini/api-prices-project/internal/platform/tests"
	"github.com/agustinrabini/api-prices-project/internal/prices"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"
	"github.com/stretchr/testify/assert"
)

const (
	priceBody = `{
		"id_item": 1,
		"amount": 10
	}`
	priceBodyNilAmount = `{
		"id_item": 1,
		"amount": 0
	}`
)

func TestPriceHandler_CreateOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "POST", "/prices", nil, priceBody)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func TestPriceHandler_CreateErrBody(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "POST", "/prices", nil, priceBodyNilAmount)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
}

func TestPriceHandler_CreateErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockRepo.HandleSave = func(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
		return nil, apierrors.NewInternalServerApiError("mock error", errors.New("mock error"))
	}
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "POST", "/prices", nil, priceBody)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestPriceHandler_GetOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	priceMock := domain.Price{Amount: 1, ItemID: "1", ID: "1"}

	mockRepo := prices.NewRepositoryMock()
	mockRepo.HandleGet = func(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
		return priceMock, nil
	}
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "GET", "/prices/1", nil, "")
	assert.Equal(t, http.StatusOK, response.Code)
	var result domain.Price
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.Equal(t, priceMock, result)
}

func TestPriceHandler_GetErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "GET", "/prices/1", nil, "")
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestPriceHandler_DeleteOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "DELETE", "/prices/1", nil, "")
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestPriceHandler_DeleteErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockRepo.HandleDelete = func(ctx context.Context, itemID string) apierrors.ApiError {
		return apierrors.NewInternalServerApiError("mock error", errors.New("mock error"))
	}
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "DELETE", "/prices/1", nil, "")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func TestPriceHandler_UpdateOk(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "PUT", "/prices/1", nil, priceBody)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestPriceHandler_UpdateErrInvalidID(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "PUT", "/prices/invalidID", nil, priceBody)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestPriceHandler_UpdateErrBody(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "PUT", "/prices/1", nil, priceBodyNilAmount)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
}

func TestPriceHandler_UpdateErr(t *testing.T) {
	teardownTestCase := tests.CustomSetupTestCase(beforeTestCase, afterTestCase)
	defer teardownTestCase()

	mockRepo := prices.NewRepositoryMock()
	mockRepo.HandleUpdate = func(ctx context.Context, priceID string, price domain.Price) apierrors.ApiError {
		return apierrors.NewInternalServerApiError("mock error", errors.New("mock error"))
	}
	mockService := prices.NewService(mockRepo)
	mockHandler := handlers.NewPricesHandler(mockService)
	depend.Prices = mockHandler

	response := executeRequest(buildRouter(), "PUT", "/prices/1", nil, priceBody)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
}
