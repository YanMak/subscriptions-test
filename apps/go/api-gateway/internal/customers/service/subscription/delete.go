package customerSubscriptionService

import (
	"context"
	"time"

	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
)

func (s *CustomerSubscriptionService) Delete(dto *customerBrokerContract.CustomerSubscriptionIDRequest) (customerBrokerContract.CustomerSubscriptionDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
	defer cancel()

	err := s.repo.DeleteWithTx(ctx, dto.ID)
	return customerBrokerContract.CustomerSubscriptionDeleteResponse{}, err
}

// func (s *CustomerSubscriptionService) ListOfSubscriptions(userID string) ([]*customerSubscriptionModel.CustomerSubscription, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
// 	defer cancel()

// 	entities, err := s.repo.ListWithTx(ctx, userID)
// 	return entities, err
// }
