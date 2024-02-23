package api

import (
	"fmt"
	"net/http"
	oc "watcher/api/controllers/onboarding"
	"watcher/api/routers"
	"watcher/service"

	"github.com/gin-gonic/gin"
)

func StartServer(services service.Services) {
	r := gin.Default()

	onboardingController := oc.NewOnboardingController(services.Onboarding)
	// eventController :=

	onboardingRoutes := r.Group("/onboarding")
	routers.SetOnboardingRoutes(onboardingRoutes, onboardingController)

	eventRoutes := r.Group("/event")
	routers.SetEventRoutes(eventRoutes)

	r.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	fmt.Println("no bullshit programmer...")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	fmt.Println("watcher running")
	r.Run(":8080")
}
