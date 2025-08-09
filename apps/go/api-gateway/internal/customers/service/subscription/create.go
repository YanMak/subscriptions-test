package customerSubscriptionService

import (
	"context"
	"time"

	"github.com/google/uuid"
	customerSubscriptionModel "project1.v0/api_gateway/internal/customers/models/subscription"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	conv "project1.v0/mappers/conv"
)

func (s *CustomerSubscriptionService) Create(dto *customerBrokerContract.CustomerSubscriptionCreateRequest) (*customerBrokerContract.CustomerSubscriptionCreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
	defer cancel()

	userID, _ := uuid.Parse(dto.UserID)

	entity := customerSubscriptionModel.NewCustomerSubscription(
		customerSubscriptionModel.NewCustomerSubscriptionDeps{
			ServiceName: dto.ServiceName,
			Price:       dto.Price,
			UserID:      userID,
			StartDate:   conv.ParseDate(dto.StartDate),
			EndDate:     conv.ParseDatePtr(dto.EndDate),
		},
	)
	id, err := s.repo.CreateWithTx(ctx, entity)
	if err != nil {
		return nil, err
	}
	return &customerBrokerContract.CustomerSubscriptionCreateResponse{ID: id}, err

}
