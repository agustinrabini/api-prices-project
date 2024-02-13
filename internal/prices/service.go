package prices

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/agustinrabini/api-prices-project/cmd/api/config"
	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"

	"github.com/agustinrabini/api-prices-project/internal/domain"
)

type Service interface {
	Get(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError)
	Create(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError)
	Update(ctx context.Context, priceID string, price domain.Price) apierrors.ApiError
	Delete(ctx context.Context, priceID string) apierrors.ApiError

	GetItemsPrices(ctx context.Context, itemsIds domain.ItemsIdsRequest) (domain.ItemsIdsResponse, apierrors.ApiError)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	if config.IsProductionEnvironment() {
		return NewServiceImpl(repository)
	}
	return NewPricesServiceMock()
}

func NewServiceImpl(repository Repository) Service {
	return &service{repo: repository}
}

func (s *service) Get(ctx context.Context, itemID string) (domain.Price, apierrors.ApiError) {
	prices, err := s.repo.Get(ctx, itemID)
	if err != nil {
		return domain.Price{}, err
	}
	return findBestPrice(prices), nil
}

func (s *service) Create(ctx context.Context, price domain.Price) (interface{}, apierrors.ApiError) {
	return s.repo.Save(ctx, price)
}

func (s *service) Update(ctx context.Context, priceID string, price domain.Price) apierrors.ApiError {
	return s.repo.Update(ctx, priceID, price)
}
func (s *service) Delete(ctx context.Context, priceID string) apierrors.ApiError {
	return s.repo.Delete(ctx, priceID)
}

func (s *service) GetItemsPrices(ctx context.Context, itemsIds domain.ItemsIdsRequest) (domain.ItemsIdsResponse, apierrors.ApiError) {

	var response domain.ItemsIdsResponse

	bb, err := json.Marshal(itemsIds)
	if err != nil {
		panic(err)
	}
	fmt.Println("la request: ", string(bb))

	for _, itemID := range itemsIds.Items {
		prices, err := s.repo.Get(ctx, itemID)
		fmt.Println("lo que trae la db: ", prices)
		if err != nil {
			return domain.ItemsIdsResponse{}, err
		}
		response.Prices = append(response.Prices, prices...)
	}

	return response, nil
}

// WARN: Si el objeto Price se vuelve muy grande usar punteros para
// no tener que andar copiando en el buffer
func findBestPrice(prices domain.Prices) domain.Price {
	if len(prices) == 0 {
		return domain.Price{}
	}
	bestPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i].Amount < bestPrice.Amount {
			bestPrice = prices[i]
		}
	}
	return bestPrice
}
