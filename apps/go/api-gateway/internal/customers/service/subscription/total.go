package customerSubscriptionService

import (
	"context"
	"time"

	"github.com/google/uuid"
	customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

func (s *CustomerSubscriptionService) TotalCost(
	dto *customerBrokerContract.CustomerSubscriptionTotalCostRequest,
) (customerBrokerContract.CustomerSubscriptionTotalCostResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000000*time.Millisecond)
	defer cancel()

	userID, _ := uuid.Parse(dto.UserID)
	entity := customerSubscriptionModel.NewCustomerSubscription(customerSubscriptionModel.NewCustomerSubscriptionDeps{
		UserID: userID,
	})
	mask := []string{}
	if dto.ServiceName != nil {
		entity.UpdateServiceName(dto.ServiceName)
		field := "ServiceName"
		mask = append(mask, field)
	}
	if dto.StartDate != nil {
		entity.UpdateStartDate(dto.StartDate)
		field := "StartDate"
		mask = append(mask, field)
	}
	if dto.EndDate != nil {
		entity.UpdateEndDate(dto.EndDate)
		field := "EndDate"
		mask = append(mask, field)
	}

	var total int
	total, err := s.repo.TotalCostWithTx(ctx, entity, mask)

	return customerBrokerContract.CustomerSubscriptionTotalCostResponse{Total: total}, err

}

// // func (s *CustomerSubscriptionService) AggregateSum(userID string) ([]*customerSubscriptionModel.CustomerSubscription, error) {
// // 	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
// // 	defer cancel()

// // 	entities, err := s.repo.AggregateSumWithTx(ctx, userID)
// // 	return entities, err
// // }
