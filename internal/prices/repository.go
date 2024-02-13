package prices

import (
	"context"
	"fmt"

	"github.com/agustinrabini/api-prices-project/cmd/api/config"
	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/matiasnu/go-jopit-toolkit/gonosql"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	errorInDB = "[%s] Error in DB"
)

var ErrPriceNotFound = apierrors.NewNotFoundApiError("price not found")

type Repository interface {
	Get(ctx context.Context, itemID string) (domain.Prices, apierrors.ApiError)
	Save(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError)
	Update(ctx context.Context, priceID string, updatePrice domain.Price) apierrors.ApiError
	Delete(ctx context.Context, priceID string) apierrors.ApiError
}

type pricesRepository struct {
	PricesCollection *mongo.Collection
}

func NewRepository(pricesCollection *mongo.Collection) Repository {
	if config.IsProductionEnvironment() {
		return NewRepositoryImpl(pricesCollection)
	}
	return NewRepositoryMock()
}

func NewRepositoryImpl(pricesCollection *mongo.Collection) Repository {
	return &pricesRepository{PricesCollection: pricesCollection}
}

func (storage *pricesRepository) Get(ctx context.Context, itemID string) (domain.Prices, apierrors.ApiError) {
	var models domain.Prices
	cursorResult, err := gonosql.GetByKey(ctx, storage.PricesCollection, "item_id", itemID)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), err)
	}
	if cursorResult.RemainingBatchLength() == 0 {
		return nil, ErrPriceNotFound
	}
	if err = cursorResult.All(ctx, &models); err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Get"), err)
	}
	return models, nil
}

func (storage *pricesRepository) Save(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
	result, err := gonosql.InsertOne(ctx, storage.PricesCollection, price)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Save"), err)
	}
	return result.InsertedID, nil
}

func (storage *pricesRepository) Update(ctx context.Context, priceID string, updatePrice domain.Price) apierrors.ApiError {
	// TODO review result for Update?
	_, err := gonosql.Update(ctx, storage.PricesCollection, priceID, updatePrice)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Update"), err)
	}
	return nil
}

func (storage *pricesRepository) Delete(ctx context.Context, priceID string) apierrors.ApiError {
	// TODO return err when ID has already been removed?
	_, err := gonosql.Delete(ctx, storage.PricesCollection, priceID)
	if err != nil {
		return apierrors.NewInternalServerApiError(fmt.Sprintf(errorInDB, "Delete"), err)
	}
	return nil
}
