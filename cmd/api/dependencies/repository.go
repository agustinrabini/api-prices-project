package dependencies

import (
	"github.com/agustinrabini/api-prices-project/internal/platform/storage"
	"github.com/agustinrabini/api-prices-project/internal/prices"
	"github.com/matiasnu/go-jopit-toolkit/gonosql"
)

const (
	KvsPricesCollection = "prices"
)

type DependencyManager struct {
	*gonosql.Data
}

func NewDependencyManager() DependencyManager {
	db := storage.NewNoSQL()
	if db.Error != nil {
		panic(db.Error)
	}
	return DependencyManager{
		db,
	}
}

func (m DependencyManager) PricesRepository() prices.Repository {
	return prices.NewRepository(m.NewCollection(KvsPricesCollection))
}
