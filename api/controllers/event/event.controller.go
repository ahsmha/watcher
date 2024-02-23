package controller

import (
	"fmt"
	"io/ioutil"
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
	default:
		fmt.Printf("in github")
		type Body struct {
			Ref    string `json:"ref"`
			Before string `json:"before"`
			After  string `json:"after"`
		}
		var b Body
		event := github.PushEventRequest{
			Source: "github",
			Type:   ctx.Request.Header.Get("x-github-event"),
		}
		var body interface{}
		bytes, err := ioutil.ReadAll(ctx.Request.Body)
		bodyStr := string(bytes)
		fmt.Printf("body: %s", bodyStr[0:10])
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{})
		// 	return
		// }
		if err := ctx.BindJSON(&b); err != nil {

			fmt.Printf("Body Unmarshalled %v", b)
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		fmt.Printf("Body Unmarshalled %v", b)
		if err := ctx.BindJSON(&event); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		err = eci.service.PushGithubEvent(ctx, &event)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{})
			return
		}
	case "gitlab":
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "event pushed to kafka successfully",
	})
}
