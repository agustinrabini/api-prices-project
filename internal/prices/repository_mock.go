package prices

import (
	"context"

	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"
)

type RepositoryMock struct {
	HandleGet    func(ctx context.Context, itemID string) (domain.Prices, apierrors.ApiError)
	HandleSave   func(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError)
	HandleUpdate func(ctx context.Context, priceID string, updatePrice domain.Price) apierrors.ApiError
	HandleDelete func(ctx context.Context, priceID string) apierrors.ApiError
}

func NewRepositoryMock() RepositoryMock {
	return RepositoryMock{}
}

func (mock RepositoryMock) Get(ctx context.Context, itemID string) (domain.Prices, apierrors.ApiError) {
	if mock.HandleGet != nil {
		return mock.HandleGet(ctx, itemID)
	}
	return nil, nil
}

func (mock RepositoryMock) Save(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
	if mock.HandleSave != nil {
		return mock.HandleSave(ctx, price)
	}
	return nil, nil
}

func (mock RepositoryMock) Update(ctx context.Context, priceID string, updatePrice domain.Price) apierrors.ApiError {
	if mock.HandleUpdate != nil {
		return mock.HandleUpdate(ctx, priceID, updatePrice)
	}
	return nil
}

func (mock RepositoryMock) Delete(ctx context.Context, priceID string) apierrors.ApiError {
	if mock.HandleDelete != nil {
		return mock.HandleDelete(ctx, priceID)
	}
	return nil
}
