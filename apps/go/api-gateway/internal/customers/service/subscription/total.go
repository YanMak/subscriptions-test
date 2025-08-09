package customerSubscriptionService

import (
	"context"
	"time"

	customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

func (s *CustomerSubscriptionService) TotalCost(
	dto *customerBrokerContract.CustomerSubscriptionTotalCostRequest,
) (customerBrokerContract.CustomerSubscriptionTotalCostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000000*time.Millisecond)
	defer cancel()

	entity := customerSubscriptionModel.NewCustomerSubscription(customerSubscriptionModel.NewCustomerSubscriptionDeps{})
	if dto.UserID != nil {
		entity.UpdateUserID(dto.UserID)
	}
	if dto.ServiceName != nil {
		entity.UpdateServiceName(dto.ServiceName)
	}
	if dto.StartDate != nil {
		entity.UpdateStartDate(dto.StartDate)
	}
	if dto.EndDate != nil {
		entity.UpdateEndDate(dto.EndDate)
	}

	var total int
	mask := s.decoder.BuildUpdateMask(dto)
	total, err := s.repo.TotalCostWithTx(ctx, entity, mask)

	return customerBrokerContract.CustomerSubscriptionTotalCostResponse{Total: total}, err

}

// // func (s *CustomerSubscriptionService) AggregateSum(userID string) ([]*customerSubscriptionModel.CustomerSubscription, error) {
// // 	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
// // 	defer cancel()

// // 	entities, err := s.repo.AggregateSumWithTx(ctx, userID)
// // 	return entities, err
// // }
