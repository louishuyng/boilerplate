package api

import (
	"net/http"
	"rz-server/internal/app/example/application"
	"rz-server/internal/common/interfaces"
)

type ExampleApi struct {
	service application.ExampleService
	util    *interfaces.Util
}

func RegisterAPI(
	server interfaces.Server,
	service application.ExampleService,
	util *interfaces.Util,
) {
	u := new(ExampleApi)
	u.service = service
	u.util = util

	server.POST("/example", u.createExample)
}

func (u *ExampleApi) createExample(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Creating example", map[string]any{
		"method": r.Method,
	})

	command := application.CreateExampleCommand{
		Name: "John",
	}

	_, _ = u.service.CreateExample(command)

	w.Write([]byte("Example created"))
	w.WriteHeader(http.StatusCreated)
}
