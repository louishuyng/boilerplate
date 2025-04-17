package example_api

import (
	"rz-server/internal/app/example/application"
	"rz-server/internal/common/interfaces"
	"rz-server/internal/common/middlewares"
)

type ExampleRouting struct {
	publicRoute  interfaces.Route
	privateRoute interfaces.Route
}

func NewRoutingSetup(
	server interfaces.Server,
	util *interfaces.Util,
	service application.ExampleService,
) *ExampleRouting {

	publicRoute := server.NewRoute()

	publicRoute.SetPathPrefix("/api")
	publicRoute.Use(middlewares.NewLoggingMiddleware(util.Log))

	privateRoute := server.NewRoute()
	privateRoute.Use(
		middlewares.NewLoggingMiddleware(util.Log),
		middlewares.NewJWTAuthorizationMiddleware(),
	)

	return &ExampleRouting{
		publicRoute:  publicRoute,
		privateRoute: privateRoute,
	}
}
