package customerSubscriptionService

import (
	"context"
	"time"

	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	conv "project1.v0/mappers/conv"
)

func (s *CustomerSubscriptionService) List(dto *customerBrokerContract.CustomerSubscriptionsListRequest) (customerBrokerContract.CustomerSubscriptionsListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
	defer cancel()

	mask := s.decoder.BuildUpdateMask(dto)
	entities, err := s.repo.ListWithTx(ctx, dto, mask)

	var list []customerBrokerContract.CustomerSubscription
	for _, entity := range entities {
		list = append(list,
			customerBrokerContract.CustomerSubscription{
				ID:          entity.ID.String(),
				ServiceName: entity.ServiceName,
				Price:       entity.Price,
				UserID:      entity.UserID.String(),
				StartDate:   entity.StartDate.Format("01-2006"),
				EndDate:     conv.FormatPtrDate(entity.EndDate, "01-2006"),
			},
		)
	}

	return customerBrokerContract.CustomerSubscriptionsListResponse{
		Subscriptions: list,
	}, err
}
