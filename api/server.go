package api

import (
	"fmt"
	"net/http"
	ec "watcher/api/controllers/event"
	"watcher/api/routers"
	"watcher/service"
	"watcher/utils"

	"github.com/gin-gonic/gin"
)

func StartServer(services service.Services, config *utils.Config) {
	r := gin.Default()

	eventController := ec.NewEventController(services.Event)

	eventRoutes := r.Group("/event")
	routers.SetEventRoutes(eventRoutes, eventController)

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	fmt.Println("no bullshit programmer...")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	r.Run(":9001")
}
