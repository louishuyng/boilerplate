package events

const (
	UserCreatedEventName = "UserCreatedEvent"
)

type UserCreatedEvent struct {
	Id   string
	Name string
}

type UserEventConsumer interface {
	StartEventLoop()
	handleUserCreatedEvent(event UserCreatedEvent) error
}
