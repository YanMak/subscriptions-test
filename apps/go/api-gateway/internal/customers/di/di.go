package di

import (
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

type ICustomerSubscriptionService interface {
	Create(*customerBrokerContract.CustomerSubscriptionCreateRequest) (*customerBrokerContract.CustomerSubscriptionCreateResponse, error)

	Get(*customerBrokerContract.CustomerSubscriptionIDRequest) (*customerBrokerContract.CustomerSubscription, error)

	Delete(*customerBrokerContract.CustomerSubscriptionIDRequest) (customerBrokerContract.CustomerSubscriptionDeleteResponse, error)

	List(*customerBrokerContract.CustomerSubscriptionsListRequest) (customerBrokerContract.CustomerSubscriptionsListResponse, error)

	Update(*customerBrokerContract.CustomerSubscriptionUpdateRequest) (customerBrokerContract.CustomerSubscriptionUpdateResponse, error)

	TotalCost(*customerBrokerContract.CustomerSubscriptionTotalCostRequest) (customerBrokerContract.CustomerSubscriptionTotalCostResponse, error)
}
