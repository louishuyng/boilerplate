package playground

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	playground_setup "rz-server/playground/setup"
	"testing"
)

func TestExamplePlayGround(t *testing.T) {
	service := playground_setup.MakeExampleService()

	_, _ = service.CreateExample(example_commands.CreateExampleCommand{
		Name: "test",
	})

	t.Log("Hello, playground!")
}
