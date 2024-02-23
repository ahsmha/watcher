package service

import (
	"watcher/domain/event"
	eventService "watcher/service/event"
)

type Services struct {
	Event event.Service
}

func InitServices() Services {
	eventService := eventService.NewEventServiceImplementation()
	return Services{
		Event: eventService,
	}
}
