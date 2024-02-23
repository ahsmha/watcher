package routers

import (
	ec "watcher/api/controllers/event"

	"github.com/gin-gonic/gin"
)

func SetEventRoutes(or *gin.RouterGroup, c ec.EventController) {
	or.POST("/", c.HandleEvent)
}
