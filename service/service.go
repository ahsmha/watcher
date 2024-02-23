package service

import (
	"watcher/domain/event"
	eventService "watcher/service/event"
	"watcher/utils"
)

type Services struct {
	Event event.Service
}

func InitServices(config *utils.Config) Services {
	eventService := eventService.NewEventServiceImplementation(config)
	return Services{
		Event: eventService,
	}
}
