package customerSubscriptionService

import (
	customerSubscriptionsRepository "project1.v0/api_gateway/internal/customers/repository/subscription"
	"project1.v0/pkg/decoder"
)

type CustomerSubscriptionServiceDeps struct {
	Repo *customerSubscriptionsRepository.CustomerSubscriptionRepository
	*decoder.DecoderService
}

type CustomerSubscriptionService struct {
	repo    *customerSubscriptionsRepository.CustomerSubscriptionRepository
	decoder *decoder.DecoderService
}

func NewCustomerSubscriptionService(deps CustomerSubscriptionServiceDeps) *CustomerSubscriptionService {
	return &CustomerSubscriptionService{repo: deps.Repo}
}
