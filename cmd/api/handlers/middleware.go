package handlers

import (
	"github.com/agustinrabini/api-prices-project/cmd/api/config"
	"github.com/gin-gonic/gin"
	"github.com/matiasnu/go-jopit-toolkit/goutils/logger"
)

func LoggerHandler(requestName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqLogger := logger.NewRequestLogger(c, requestName, config.LogRatio, config.LogBodyRatio)
		c.Next()
		reqLogger.LogResponse(c)
	}
}
