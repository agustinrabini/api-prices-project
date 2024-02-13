package dependencies

import (
	"github.com/agustinrabini/api-prices-project/cmd/api/handlers"
	"github.com/agustinrabini/api-prices-project/internal/prices"
)

type Dependencies interface {
	PricesRepository() prices.Repository
}

func GetDependencyManager() Dependencies {
	return NewDependencyManager()
}

func BuildDependencies() (HandlersStruct, error) {
	depManager := GetDependencyManager()

	// Repository
	pricesRepository := depManager.PricesRepository()

	// Services
	pricesService := prices.NewService(pricesRepository)

	// Handlers
	pricesHandler := handlers.NewPricesHandler(pricesService)

	handler := HandlersStruct{
		Prices: pricesHandler,
	}
	return handler, nil
}

type HandlersStruct struct {
	Prices handlers.PricesHandler
}
