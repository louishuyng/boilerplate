package events

type EventName string

const (
	UserCreatedEventName EventName = "UserCreatedEvent"
)

type Event struct {
	Name EventName
	Body struct{}
}

type UserCreatedEvent struct {
	Id   string
	Name string
}

type UserEventConsumer interface {
	EventLoop()
	HandleUserCreatedEvent(event UserCreatedEvent) error
}
