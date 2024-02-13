package tests

import (
	"github.com/matiasnu/go-jopit-toolkit/goutils/logger"
	"github.com/matiasnu/go-jopit-toolkit/rest"
	"net/http"
)

const defaultURLBase = "http://test.jopit.com"

func mockServerConfig(statusCode int, url, method, respBody string, respHeaders http.Header, expectedCallCount int) *rest.Mock {
	mockServerConfig := new(rest.Mock)
	mockServerConfig.URL = defaultURLBase + url
	mockServerConfig.HTTPMethod = method
	mockServerConfig.RespHTTPCode = statusCode
	mockServerConfig.RespBody = respBody
	mockServerConfig.RespHeaders = respHeaders
	mockServerConfig.ExpectedCallCount = expectedCallCount
	return mockServerConfig
}

func AddMockupsServer(statusCode int, url, method, respBody string, respHeaders http.Header, optionals ...int) {
	expectedCallCount := 1
	if len(optionals) > 0 {
		expectedCallCount = optionals[0]
	}
	mockServerConfig := mockServerConfig(statusCode, url, method, respBody, respHeaders, expectedCallCount)
	err := rest.AddMockups(mockServerConfig)
	if err != nil {
		logger.Panic("Error configuring mockup server", err)
	}
}
