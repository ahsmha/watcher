package controller

import (
	"fmt"
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
		if strings.Contains(ua, val) {
			ua = val
			break
		}
	}

	switch ua {
	case "github":
		fmt.Println("Inside github")
		event := github.PushEventRequest{
			Source: "github",
			Type:   ctx.Request.Header.Get("x-github-event"),
		}
		type Body struct {
			Ref    string `json:"ref"`
			Before string `json:"before"`
			After  string `json:"after"`
		}
		var body Body
		if err := ctx.ShouldBindJSON(&body); err != nil {
			fmt.Printf("body: %v", body)
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		fmt.Printf("body: %v", body)
		err := eci.service.PushGithubEvent(ctx, &event)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{})
			return
		}
	case "gitlab":
	default:
	}

	fmt.Println("message: event pushed to kafka successfully")
}
