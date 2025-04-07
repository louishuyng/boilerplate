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

type UserEventProducer interface {
	ProduceUserCreatedEvent(event UserCreatedEvent) error
}
