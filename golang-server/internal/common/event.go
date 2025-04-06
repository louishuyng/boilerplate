package common

type Event struct {
	Name string
	Body struct{}
}

func NewEventChannel() chan Event {
	return make(chan Event, 100)
}
