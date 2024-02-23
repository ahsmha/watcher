package service

import (
	"watcher/domain/onboarding"
	onboardingService "watcher/service/onboarding"
)

type Services struct {
	Onboarding onboarding.Service
	// Event      event.Service
}

func InitServices() Services {
	onboardingService := onboardingService.NewOnboardingServiceImplementation()
	return Services{
		Onboarding: onboardingService,
	}
}
