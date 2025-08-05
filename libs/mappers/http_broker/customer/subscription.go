package customerHttpToBrokerMapper

import (
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	customerHttpContract "project1.v0/contracts/transport/http/customer"
)

func CustomerSubscriptionCreateRequest_HttpToBroker(dto *customerHttpContract.CustomerSubscriptionCreateRequest) *customerBrokerContract.CustomerSubscriptionCreateRequest {
	return &customerBrokerContract.CustomerSubscriptionCreateRequest{
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserID:      dto.UserID,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}

func CustomerSubscriptionUpdateRequest_HttpToBroker(
	id string,
	dto *customerHttpContract.CustomerSubscriptionUpdateRequest,
) *customerBrokerContract.CustomerSubscriptionUpdateRequest {
	return &customerBrokerContract.CustomerSubscriptionUpdateRequest{
		ID:          id,
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}

func CustomerSubscriptionIDRequest_HttpToBroker(
	dto *customerHttpContract.CustomerSubscriptionIDPath,
) *customerBrokerContract.CustomerSubscriptionIDRequest {
	return &customerBrokerContract.CustomerSubscriptionIDRequest{
		ID: *dto.ID,
	}
}

func CustomerSubscriptionListRequest_HttpToBroker(dto *customerHttpContract.SubscriptionListQuery) *customerBrokerContract.CustomerSubscriptionsListRequest {
	return &customerBrokerContract.CustomerSubscriptionsListRequest{
		Limit:       dto.Limit,
		Offset:      dto.Offset,
		ServiceName: dto.ServiceName,
		//Price:       dto.Price,
		UserID:    dto.UserID,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
	}
}

func CustomerSubscriptionTotalCostRequest_HttpToBroker(
	dto *customerHttpContract.SubscriptionTotalQuery,
) *customerBrokerContract.CustomerSubscriptionTotalCostRequest {
	return &customerBrokerContract.CustomerSubscriptionTotalCostRequest{
		UserID:      dto.UserID,
		ServiceName: dto.ServiceName,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}
}
