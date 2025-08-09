package customerSubscriptionService

import (
	"context"
	"time"

	"github.com/google/uuid"
	customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

func (s *CustomerSubscriptionService) Update(dto *customerBrokerContract.CustomerSubscriptionUpdateRequest) (customerBrokerContract.CustomerSubscriptionUpdateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000000*time.Millisecond)
	defer cancel()

	id, _ := uuid.Parse(dto.ID)

	// we do not retrieve entity from db previously
	// for performance reason
	entity := customerSubscriptionModel.NewCustomerSubscription(customerSubscriptionModel.NewCustomerSubscriptionDeps{
		ID: id,
	})
	//mask := []string{}
	if dto.ServiceName != nil {
		entity.UpdateServiceName(dto.ServiceName)
		// field := "ServiceName"
		// mask = append(mask, field)
	}
	if dto.Price != nil {
		entity.UpdatePrice(dto.Price)
		// field := "Price"
		// mask = append(mask, field)
	}
	if dto.StartDate != nil {
		entity.UpdateStartDate(dto.StartDate)
		// field := "StartDate"
		// mask = append(mask, field)
	}
	if dto.EndDate != nil {
		entity.UpdateEndDate(dto.EndDate)
		// field := "EndDate"
		// mask = append(mask, field)
	}

	mask := s.decoder.BuildUpdateMask(dto)
	err := s.repo.UpdateWithTx(ctx, entity, mask)

	return customerBrokerContract.CustomerSubscriptionUpdateResponse{}, err
}
