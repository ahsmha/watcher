package controller

import (
	"net/http"
	"strings"
	"watcher/domain/event"
	"watcher/entities/events/github"
	"watcher/utils"

	"github.com/gin-gonic/gin"
)

type EventController interface {
	HandleEvent(ctx *gin.Context)
}

type eventControllerImplementation struct {
	service event.Service
}

func NewEventController(service event.Service) EventController {
	return &eventControllerImplementation{
		service: service,
	}
}

func (eci *eventControllerImplementation) HandleEvent(ctx *gin.Context) {
	// check if webhook repo has push webhook active
	// if (isHookActive()) {
	ua := strings.ToLower(ctx.Request.Header.Get("User-Agent"))

	for _, val := range utils.HookSource {
		if strings.Contains(val, ua) {
			ua = val
			break
		}
	}

	switch ua {
	case "github":
		event := github.PushEventRequest{
			Source: "github",
			Type:   ctx.Request.Header.Get("x-github-event"),
		}
		if err := ctx.ShouldBindJSON(&event); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		err := eci.service.PushGithubEvent(ctx, &event)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{})
			return
		}
	case "gitlab":
	default:
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event pushed to kafka successfully",
	})
}
