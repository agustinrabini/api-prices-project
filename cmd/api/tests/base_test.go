package tests

import (
	"os"
	"testing"

	"github.com/agustinrabini/api-prices-project/cmd/api/app"
	"github.com/agustinrabini/api-prices-project/cmd/api/dependencies"
	"github.com/gin-gonic/gin"
	"github.com/matiasnu/go-jopit-toolkit/rest"
)

var depend dependencies.HandlersStruct

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// Helper function to create a router during testing.
func buildRouter() *gin.Engine {
	router := app.ConfigureRouter()
	app.MapUrlsToControllers(router, depend)
	return router
}

func setup() {
	if err := os.Setenv("IS_PROD_SCOPE", "false"); err != nil {
		panic(err)
	}
}

func teardown() {
}

func beforeTestCase() {
	depend, _ = dependencies.BuildDependencies()
}

func afterTestCase() {
	rest.FlushMockups()
}
