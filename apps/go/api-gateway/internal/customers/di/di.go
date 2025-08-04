package di

import (
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

type ICustomerSubscriptionService interface {
	Create(dto *customerBrokerContract.CustomerSubscriptionCreateRequest) (*customerBrokerContract.CustomerSubscriptionCreateResponse, error)

	// Get(customerBrokerContract.CustomerSubscriptionGetRequest) (*customerBrokerContract.CustomerSubscription, error)

	// List(dto customerBrokerContract.CustomerSubscriptionsListRequest) (customerBrokerContract.CustomerSubscriptionsListResponse, error)

	// Delete(customerBrokerContract.CustomerSubscriptionDeleteRequest) (customerBrokerContract.CustomerSubscriptionDeleteResponse, error)

	Update(dto *customerBrokerContract.CustomerSubscriptionUpdateRequest) (customerBrokerContract.CustomerSubscriptionUpdateResponse, error)

	// TotalCost(dto *customerBrokerContract.CustomerSubscriptionTotalCostRequest) (customerBrokerContract.CustomerSubscriptionTotalCostResponse, error)
}
