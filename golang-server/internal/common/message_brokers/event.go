package message_brokers

import "rz-server/internal/common/interfaces"

func NewEventChannel() chan interfaces.Event {
	return make(chan interfaces.Event, 100)
}
