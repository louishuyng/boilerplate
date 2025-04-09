package events

import (
	"rz-server/internal/app/example/application"
	"rz-server/internal/common/interfaces"
	"time"
)

var _ ExampleEventConsumer = &InMemoryExampleEventConsumer{}

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
			case ExampleCreatedEventName:
				// Parse event to ExampleCreatedEvent

				exampleCreatedEvent := ExampleCreatedEvent{}

				_ = inMemory.HandleExampleCreatedEvent(exampleCreatedEvent)
			}

		case <-time.After(1 * time.Second):
		}
	}
}

func (inMemory *InMemoryExampleEventConsumer) HandleExampleCreatedEvent(event ExampleCreatedEvent) error {
	command := application.CreateExampleCommand{
		Name: event.Name,
	}

	_, _ = inMemory.application.CreateExample(command)

	return nil
}
