package example_api

import (
	"net/http"
	"rz-server/internal/app/example/application"
	example_commands "rz-server/internal/app/example/application/example/commands"
	"rz-server/internal/common/interfaces"
)

type ExampleApi struct {
	service application.ExampleService
	util    *interfaces.Util
	server  interfaces.Server
}

func New(
	server interfaces.Server,
	service application.ExampleService,
	util *interfaces.Util,
) *ExampleApi {
	u := new(ExampleApi)
	u.service = service
	u.util = util
	u.server = server

	return u
}

func (u *ExampleApi) Register() {
	u.server.POST("/example", u.createExample)
}

func (u *ExampleApi) createExample(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Creating example", map[string]any{
		"method": r.Method,
	})

	command := example_commands.CreateExampleCommand{
		Name: "John",
	}

	_, err := u.service.CreateExample(command)

	if err != nil {
		w.Write([]byte(err.GetMessage()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Example created"))
	w.WriteHeader(http.StatusCreated)
}
