package app

import (
	"fmt"
	"time"

	"github.com/agustinrabini/api-prices-project/cmd/api/config"
	"github.com/agustinrabini/api-prices-project/cmd/api/dependencies"

	"github.com/gin-gonic/gin"
	"github.com/matiasnu/go-jopit-toolkit/gingonic/handlers"
	"github.com/matiasnu/go-jopit-toolkit/goutils/logger"
)

func Start() {
	handler, errBuildDepend := dependencies.BuildDependencies()
	if errBuildDepend != nil {
		fmt.Printf("Error Build Dependencies")
		waitAndPanic(errBuildDepend)
	}
	jopitRouter := ConfigureRouter()
	MapUrlsToControllers(jopitRouter, handler)

	if errRouter := jopitRouter.Run(config.ConfMap.APIRestServerPort); errRouter != nil {
		logger.Errorf("Error starting router", errRouter)
		waitAndPanic(errRouter)
	}
}

func ConfigureRouter() *gin.Engine {
	logger.InitLog(config.ConfMap.LoggingPath, config.ConfMap.LoggingFile, config.ConfMap.LoggingLevel)
	return handlers.CustomJopitRouter(handlers.JopitRouterConfig{DisableFirebaseAuth: true})
}

func waitAndPanic(err error) {
	time.Sleep(2 * time.Second) // needs one second to send the log
	panic(err)
}
