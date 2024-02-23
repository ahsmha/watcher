package eventService

import "watcher/domain/event"

type eventServiceImplementation struct {
}

func NewEventServiceImplementation() event.Service {
	return &eventServiceImplementation{}
}

func (osi *eventServiceImplementation) GetLineCount() int {
	return 0
}
