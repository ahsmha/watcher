package event

import (
	"watcher/entities/events/github"

	"github.com/gin-gonic/gin"
)

type Service interface {
	PushGithubEvent(ctx *gin.Context, event *github.PushEventRequest) error
}
