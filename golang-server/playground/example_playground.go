package main

import (
	"fmt"
	example_commands "rz-server/internal/app/example/application/example/commands"
	playground_setup "rz-server/playground/setup"
)

func main() {
	service := playground_setup.MakeExampleService()

	_, _ = service.CreateExample(example_commands.CreateExampleCommand{
		Name: "test",
	})

	fmt.Println("Hi")
}
