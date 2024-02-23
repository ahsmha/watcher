package controller

import (
	"fmt"
	"watcher/domain/onboarding"

	"github.com/gin-gonic/gin"
)

type OnboardingController interface {
	HandleOnboarding(ctx *gin.Context)
}

type onboardingControllerImplementation struct {
	service onboarding.Service
}

func NewOnboardingController(service onboarding.Service) OnboardingController {
	return &onboardingControllerImplementation{
		service: service,
	}
}

func (oci *onboardingControllerImplementation) HandleOnboarding(ctx *gin.Context) {
	xyz := oci.service.GetLineCount()
	fmt.Println(xyz)
}
