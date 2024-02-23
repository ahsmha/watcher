package routers

import (
	oc "watcher/api/controllers/onboarding"

	"github.com/gin-gonic/gin"
)

func SetOnboardingRoutes(or *gin.RouterGroup, c oc.OnboardingController) {
	or.POST("/", c.HandleOnboarding)
}
