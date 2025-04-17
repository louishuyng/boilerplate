package example_api

import (
	"net/http"
	example_resources "rz-server/internal/app/example/api/example/resources"
	"rz-server/internal/app/example/application"
	example_commands "rz-server/internal/app/example/application/example/commands"
	json_helper "rz-server/internal/common/helpers/json"
	"rz-server/internal/common/interfaces"
)

type ExampleApi struct {
	service application.ExampleService
	util    *interfaces.Util
	*ExampleRouting
}

func New(
	server interfaces.Server,
	service application.ExampleService,
	util *interfaces.Util,
) *ExampleApi {
	u := new(ExampleApi)
	u.service = service
	u.util = util
	u.ExampleRouting = NewRoutingSetup(server, util, service)

	return u
}

func (u *ExampleApi) Register() {
	u.publicRoute.Get("/example", u.createExample)
}

func (u *ExampleApi) createExample(w http.ResponseWriter, r *http.Request) {
	u.util.Log.Info("Creating example", map[string]any{
		"method": r.Method,
	})

	command := example_commands.CreateExampleCommand{
		Name: "John",
	}

	data, err := u.service.CreateExample(command)

	if err != nil {
		json_helper.RespondJsonError(err, w)
		return
	}

	json_helper.RespondJsonResourceSuccess(
		example_resources.NewCreateExampleMapper(data),
		w,
	)
}
