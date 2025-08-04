package customerSubscriptionService

import customerSubscriptionsRepository "project1.v0/api_gateway/internal/customers/repository/subscription"

type CustomerSubscriptionServiceDeps struct {
	Repo *customerSubscriptionsRepository.CustomerSubscriptionRepository
}

type CustomerSubscriptionService struct {
	repo *customerSubscriptionsRepository.CustomerSubscriptionRepository
}

func NewCustomerSubscriptionService(deps CustomerSubscriptionServiceDeps) *CustomerSubscriptionService {
	return &CustomerSubscriptionService{repo: deps.Repo}
}
