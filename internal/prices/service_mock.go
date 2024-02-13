package prices

import (
	"context"

	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"
)

type ServiceMock struct {
	HandleGet    func(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError)
	HandleDelete func(ctx context.Context, priceID string) apierrors.ApiError
	HandleCreate func(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError)
	HandleUpdate func(ctx context.Context, priceID string, updateItem domain.Price) apierrors.ApiError

	HandleGetItemsPrices func(ctx context.Context, itemsIds domain.ItemsIdsRequest) (domain.ItemsIdsResponse, apierrors.ApiError)
}

func NewPricesServiceMock() ServiceMock {
	return ServiceMock{}
}

func (mock ServiceMock) Get(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
	if mock.HandleGet != nil {
		return mock.HandleGet(ctx, itemID)
	}
	return domain.Price{}, nil
}

func (mock ServiceMock) Delete(ctx context.Context, priceID string) apierrors.ApiError {
	if mock.HandleDelete != nil {
		return mock.HandleDelete(ctx, priceID)
	}
	return nil
}

func (mock ServiceMock) Create(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
	if mock.HandleCreate != nil {
		return mock.HandleCreate(ctx, price)
	}
	return nil, nil
}

func (mock ServiceMock) Update(ctx context.Context, priceID string, updatePrice domain.Price) apierrors.ApiError {
	if mock.HandleUpdate != nil {
		return mock.HandleUpdate(ctx, priceID, updatePrice)
	}
	return nil
}

func (mock ServiceMock) GetItemsPrices(ctx context.Context, itemsIds domain.ItemsIdsRequest) (domain.ItemsIdsResponse, apierrors.ApiError) {
	if mock.HandleGetItemsPrices != nil {
		return mock.HandleGetItemsPrices(ctx, itemsIds)
	}
	return domain.ItemsIdsResponse{}, nil
}
