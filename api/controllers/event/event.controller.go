package controller

import (
	"fmt"
	"log"
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

type Common struct {
	service event.Service
}

func NewEventController(service event.Service) EventController {
	return &Common{
		service: service,
	}
}

func (com *Common) HandleEvent(ctx *gin.Context) {
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
		log.Printf("Github Hook Request")
		var body github.GithubHookRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		log.Printf("body: %v", body)
		err := com.service.PushGithubEvent(ctx, &body)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "hook published"})
	case "gitlab":
	default:
	}

	fmt.Println("message: event pushed to kafka successfully")
}
