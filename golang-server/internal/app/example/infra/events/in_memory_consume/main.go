package in_memory_consume_event

import (
	"rz-server/internal/app/example/application"
	example_commands "rz-server/internal/app/example/application/example/commands"
	"rz-server/internal/app/example/infra/events"
	"rz-server/internal/common/interfaces"
	"time"
)

var _ events.ExampleEventConsumer = &InMemoryExampleEventConsumer{}

type InMemoryExampleEventConsumer struct {
	event       <-chan interfaces.Event
	application application.ExampleService
}

func NewInMemoryExampleEventConsumer(
	event <-chan interfaces.Event,
	application application.ExampleService,
) *InMemoryExampleEventConsumer {
	return &InMemoryExampleEventConsumer{
		event:       event,
		application: application,
	}
}

func (inMemory *InMemoryExampleEventConsumer) StartEventLoop() {
	for {
		select {
		case event := <-inMemory.event:
			switch event.Name {
			case events.ExampleCreatedEventName:
				// Parse event to ExampleCreatedEvent

				exampleCreatedEvent := events.ExampleCreatedEvent{}

				_ = inMemory.HandleExampleCreatedEvent(exampleCreatedEvent)
			}

		case <-time.After(1 * time.Second):
		}
	}
}

func (inMemory *InMemoryExampleEventConsumer) HandleExampleCreatedEvent(event events.ExampleCreatedEvent) error {
	command := example_commands.CreateExampleCommand{
		Name: event.Name,
	}

	_, _ = inMemory.application.CreateExample(command)

	return nil
}
