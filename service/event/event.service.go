package eventService

import (
	"fmt"
	"watcher/domain/event"
	"watcher/entities/events/github"
	"watcher/utils"

	"github.com/gin-gonic/gin"
)

type eventServiceImplementation struct {
	token string
}

func NewEventServiceImplementation(config *utils.Config) event.Service {
	return &eventServiceImplementation{
		token: config.Token,
	}
}

func (osi *eventServiceImplementation) PushGithubEvent(ctx *gin.Context, event *github.PushEventRequest) error {
	// push to kafka with proper payload
	fmt.Printf("%v", event)
	return nil
}
