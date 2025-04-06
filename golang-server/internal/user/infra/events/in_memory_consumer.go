package events

import (
	"go-server/internal/common"
	"go-server/internal/user/application"
	"time"
)

var _ UserEventConsumer = &InMemoryUserEventConsumer{}

type InMemoryUserEventConsumer struct {
	events      <-chan common.Event
	application application.UserService
}

func NewInMemoryUserEventConsumer(
	events <-chan common.Event,
	application application.UserService,
) *InMemoryUserEventConsumer {
	return &InMemoryUserEventConsumer{
		events:      events,
		application: application,
	}
}

func (inMemory *InMemoryUserEventConsumer) StartEventLoop() {
	for {
		for event := range inMemory.events {
			switch event.Name {
			case UserCreatedEventName:
				// Parse event to UserCreatedEvent

				userCreatedEvent := UserCreatedEvent{}

				_ = inMemory.handleUserCreatedEvent(userCreatedEvent)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (inMemory *InMemoryUserEventConsumer) handleUserCreatedEvent(event UserCreatedEvent) error {
	command := application.CreateUserCommand{
		Name: event.Name,
	}

	_, _ = inMemory.application.CreateUser(command)

	return nil
}
