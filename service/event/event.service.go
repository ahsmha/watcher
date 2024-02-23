package eventService

import (
	"watcher/domain/event"
	"watcher/utils"
)

type eventServiceImplementation struct {
	token string
}

func NewEventServiceImplementation(config *utils.Config) event.Service {
	return &eventServiceImplementation{
		token: config.Token,
	}
}

func (osi *eventServiceImplementation) Push() int {
	return 0
}
