package customerSubscriptionService

import (
	"context"
	"time"

	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	conv "project1.v0/mappers/conv"
)

func (s *CustomerSubscriptionService) Get(dto *customerBrokerContract.CustomerSubscriptionIDRequest) (*customerBrokerContract.CustomerSubscription, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
	defer cancel()

	entity, err := s.repo.GetWithTx(ctx, dto.ID)
	if err != nil {
		return nil, err
	}

	brokerRes := &customerBrokerContract.CustomerSubscription{
		ID:          entity.ID.String(),
		ServiceName: entity.ServiceName,
		Price:       entity.Price,
		UserID:      entity.UserID.String(),
		StartDate:   entity.StartDate.Format("01-2006"),
		EndDate:     conv.FormatPtrDate(entity.EndDate, "01-2006"),
	}
	// if entity.EndDate != nil {
	// 	brokerRes.EndDate = new(string)
	// 	*brokerRes.EndDate = entity.EndDate.Format("01-2006")
	// }
	return brokerRes, err
}
