package app

import (
	"github.com/agustinrabini/api-prices-project/cmd/api/dependencies"
	"github.com/agustinrabini/api-prices-project/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func MapUrlsToControllers(router *gin.Engine, h dependencies.HandlersStruct) {
	// Health
	health := handlers.NewHealthCheckerHandler()
	router.GET("/ping", health.Ping)

	router.GET("/prices/item/:id", handlers.LoggerHandler("GetPrice"), h.Prices.Get)
	router.POST("/prices", handlers.LoggerHandler("CreatePrice"), h.Prices.Create)
	router.PUT("/prices/:id", handlers.LoggerHandler("UpdatePrice"), h.Prices.Update)
	router.DELETE("/prices/:id", handlers.LoggerHandler("DeletePrice"), h.Prices.Delete)

	router.POST("/prices/items", handlers.LoggerHandler("GetItemsPrices"), h.Prices.GetItemsPrices)
}
