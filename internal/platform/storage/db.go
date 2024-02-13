package storage

import (
	jopitConfig "github.com/agustinrabini/api-prices-project/cmd/api/config"
	"github.com/matiasnu/go-jopit-toolkit/gonosql"
)

func NewNoSQL() *gonosql.Data {
	config := getDBConfig()
	return gonosql.NewNoSQL(config)
}

func getDBConfig() gonosql.Config {
	return gonosql.Config{
		Username: jopitConfig.ConfMap.MongoUser,
		Password: jopitConfig.ConfMap.MongoPassword,
		Host:     jopitConfig.ConfMap.MongoHost,
		Database: jopitConfig.ConfMap.MongoDataBase,
	}
}
