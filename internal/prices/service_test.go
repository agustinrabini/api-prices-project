package prices

import (
	"context"
	"errors"
	"testing"

	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"
	"github.com/stretchr/testify/assert"
)

const (
	mockServiceErr = "mock service err"
)

var (
	testPrice = domain.Price{
		ID:     "1",
		ItemID: "1",
		Amount: 10,
	}
)

func TestService_NewService(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	assert.NotNil(t, serviceMock)
}

func TestService_GetOk(t *testing.T) {
	repoMock := NewRepositoryMock()
	repoMock.HandleGet = func(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
		return testPrice, nil
	}
	serviceMock := NewService(repoMock)
	result, err := serviceMock.Get(nil, "1")
	assert.Nil(t, err)
	assert.EqualValues(t, testPrice, result)
}

func TestService_GetErr(t *testing.T) {
	repoMock := NewRepositoryMock()
	repoMock.HandleGet = func(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
		return domain.Price{}, apierrors.NewInternalServerApiError(mockServiceErr, errors.New(mockServiceErr))
	}
	serviceMock := NewService(repoMock)
	_, err := serviceMock.Get(nil, "1")
	assert.EqualValues(t, mockServiceErr, err.Error())
}

func TestService_CreateOk(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	// TODO check insertedID
	_, err := serviceMock.Create(nil, domain.Price{Amount: 1})
	assert.Nil(t, err)
}

func TestService_CreateErrAmount(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	_, err := serviceMock.Create(nil, domain.Price{})
	assert.EqualValues(t, "amount cant be negative or 0. Object not created", err.Error())
}

func TestService_CreateErrDB(t *testing.T) {
	repoMock := NewRepositoryMock()
	repoMock.HandleSave = func(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
		return nil, apierrors.NewInternalServerApiError(mockServiceErr, errors.New(mockServiceErr))
	}
	serviceMock := NewService(repoMock)
	_, err := serviceMock.Create(nil, domain.Price{Amount: 1})
	assert.EqualValues(t, mockServiceErr, err.Error())
}

func TestService_UpdateOk(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	err := serviceMock.Update(nil, "1", domain.Price{Amount: 1})
	assert.Nil(t, err)
}

func TestService_UpdateErrAmount(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	err := serviceMock.Update(nil, "1", domain.Price{})
	assert.EqualValues(t, "amount cant be negative or 0. Object not updated", err.Error())
}

func TestService_UpdateErrDB(t *testing.T) {
	repoMock := NewRepositoryMock()
	repoMock.HandleUpdate = func(ctx context.Context, priceID string, price domain.Price) apierrors.ApiError {
		return apierrors.NewInternalServerApiError(mockServiceErr, errors.New(mockServiceErr))
	}
	serviceMock := NewService(repoMock)
	err := serviceMock.Update(nil, "1", domain.Price{Amount: 1})
	assert.EqualValues(t, mockServiceErr, err.Error())
}

func TestService_DeleteOk(t *testing.T) {
	repoMock := NewRepositoryMock()
	serviceMock := NewService(repoMock)
	err := serviceMock.Delete(nil, "1")
	assert.Nil(t, err)
}

func TestService_DeleteErr(t *testing.T) {
	repoMock := NewRepositoryMock()
	repoMock.HandleDelete = func(ctx context.Context, priceID string) apierrors.ApiError {
		return apierrors.NewInternalServerApiError(mockServiceErr, errors.New(mockServiceErr))
	}
	serviceMock := NewService(repoMock)
	err := serviceMock.Delete(nil, "1")
	assert.EqualValues(t, mockServiceErr, err.Error())
}
