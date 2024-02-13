package tests

import "github.com/matiasnu/go-jopit-toolkit/rest"

func SetupTestCase() func() {
	rest.StartMockupServer()
	return func() {
		rest.ValidateCallCounts()
		rest.StopMockupServer()
	}
}

func CustomSetupTestCase(before func(), after func()) func() {
	rest.StartMockupServer()
	before()
	return func() {
		after()
		rest.ValidateCallCounts()
		rest.StopMockupServer()
	}
}
