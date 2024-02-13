package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/matiasnu/go-jopit-toolkit/goutils/logger"

	"github.com/matiasnu/go-jopit-toolkit/goutils/apierrors"

	"github.com/agustinrabini/api-prices-project/internal/domain"
	"github.com/agustinrabini/api-prices-project/internal/prices"
)

const (
	genericErrorMessageDecoder = "Error decoder in price handler, invalid JSON body"
)

type PricesHandler struct {
	Service prices.Service
}

func NewPricesHandler(service prices.Service) PricesHandler {
	return PricesHandler{Service: service}
}

// Get godoc
// @Summary Get price detail
// @Description Get details of price by item ID
// @Tags Prices
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} domain.Price
// @Router /prices/item/{id} [get]
func (h PricesHandler) Get(c *gin.Context) {
	itemID := c.Param("id")
	price, err := h.Service.Get(c, itemID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, price)
}

// Create godoc
// @Summary Create Price
// @Description Create price in db
// @Tags Prices
// @Accept  json
// @Produce  json
// @Param price body domain.Price true "Add price"
// @Success 200
// @Router /prices [post]
func (h PricesHandler) Create(c *gin.Context) {
	var priceInput domain.Price
	if err := binding.JSON.Bind(c.Request, &priceInput); err != nil {
		logger.Error(genericErrorMessageDecoder, err)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError(genericErrorMessageDecoder+err.Error()))
		return
	}
	insertedID, err := h.Service.Create(c, priceInput)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, insertedID)
}

// Update godoc
// @Summary Update Price
// @Description Update price in db
// @Tags Prices
// @Accept  json
// @Produce  json
// @Param id path string true "Price ID"
// @Param price body domain.Price true "Add price"
// @Success 200
// @Router /prices/{id} [put]
func (h PricesHandler) Update(c *gin.Context) {
	priceID := c.Param("id")
	var priceInput domain.Price
	if err := binding.JSON.Bind(c.Request, &priceInput); err != nil {
		logger.Error(genericErrorMessageDecoder, err)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError(genericErrorMessageDecoder+err.Error()))
		return
	}
	err := h.Service.Update(c, priceID, priceInput)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.Status(http.StatusOK)
}

// Delete godoc
// @Summary Delete Price
// @Description Delete price in db
// @Tags Prices
// @Accept  json
// @Produce  json
// @Param id path string true "Price ID"
// @Success 200
// @Router /prices/{id} [delete]
func (h PricesHandler) Delete(c *gin.Context) {
	priceID := c.Param("id")
	err := h.Service.Delete(c, priceID)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h PricesHandler) GetItemsPrices(c *gin.Context) {

	var priceInput domain.ItemsIdsRequest

	if err := binding.JSON.Bind(c.Request, &priceInput); err != nil {
		logger.Error(genericErrorMessageDecoder, err)
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError(genericErrorMessageDecoder+err.Error()))
		return
	}

	itemsIdResponse, err := h.Service.GetItemsPrices(c, priceInput)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, itemsIdResponse)
}
