package onboardingService

import "watcher/domain/onboarding"

type onboardingServiceImplementation struct {
}

func NewOnboardingServiceImplementation() onboarding.Service {
	return &onboardingServiceImplementation{}
}

func (osi *onboardingServiceImplementation) GetLineCount() int {
	return 0
}
