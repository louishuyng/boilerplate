package events

import (
	"go-server/internal/user/application"
	"time"
)

var _ UserEventConsumer = &InMemoryUserEventConsumer{}

type InMemoryUserEventConsumer struct {
	events      chan Event
	application application.UserService
}

func NewInMemoryUserEventConsumer(
	application application.UserService,
) *InMemoryUserEventConsumer {
	return &InMemoryUserEventConsumer{
		events:      make(chan Event, 100),
		application: application,
	}
}

func (inMemory *InMemoryUserEventConsumer) HandleUserCreatedEvent(event UserCreatedEvent) error {
	command := application.CreateUserCommand{
		Name: event.Name,
	}

	_, _ = inMemory.application.CreateUser(command)

	return nil
}

func (inMemory *InMemoryUserEventConsumer) EventLoop() {
	for {
		for event := range inMemory.events {
			switch event.Name {
			case UserCreatedEventName:
				// Parse event to UserCreatedEvent

				userCreatedEvent := UserCreatedEvent{}

				_ = inMemory.HandleUserCreatedEvent(userCreatedEvent)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
