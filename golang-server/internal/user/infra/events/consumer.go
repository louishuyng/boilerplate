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
	HandleUserCreatedEvent(event UserCreatedEvent) error
}
