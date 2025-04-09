package events

const (
	ExampleCreatedEventName = "ExampleCreatedEvent"
)

type ExampleCreatedEvent struct {
	Id   string
	Name string
}

type ExampleEventConsumer interface {
	StartEventLoop()
	HandleExampleCreatedEvent(event ExampleCreatedEvent) error
}

type ExampleEventProducer interface {
	ProduceExampleCreatedEvent(event ExampleCreatedEvent) error
}
