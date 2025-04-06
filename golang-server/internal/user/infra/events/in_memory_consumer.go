package events

import (
	"go-server/internal/common"
	"go-server/internal/user/application"
	"time"
)

var _ UserEventConsumer = &InMemoryUserEventConsumer{}

type InMemoryUserEventConsumer struct {
	event       <-chan common.Event
	application application.UserService
}

func NewInMemoryUserEventConsumer(
	event <-chan common.Event,
	application application.UserService,
) *InMemoryUserEventConsumer {
	return &InMemoryUserEventConsumer{
		event:       event,
		application: application,
	}
}

func (inMemory *InMemoryUserEventConsumer) StartEventLoop() {
	for {
		select {
		case event := <-inMemory.event:
			switch event.Name {
			case UserCreatedEventName:
				// Parse event to UserCreatedEvent

				userCreatedEvent := UserCreatedEvent{}

				_ = inMemory.HandleUserCreatedEvent(userCreatedEvent)
			}

		case <-time.After(1 * time.Second):
		}
	}
}

func (inMemory *InMemoryUserEventConsumer) HandleUserCreatedEvent(event UserCreatedEvent) error {
	command := application.CreateUserCommand{
		Name: event.Name,
	}

	_, _ = inMemory.application.CreateUser(command)

	return nil
}
