package routers

import (
	ec "watcher/api/controllers/event"

	"github.com/gin-gonic/gin"
)

func SetEventRoutes(er *gin.RouterGroup, c ec.EventController) {
	er.POST("/", c.HandleEvent)
}
