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

// func (s *CustomerSubscriptionService) Get(dto customerBrokerContract.CustomerSubscriptionGetRequest) (*customerBrokerContract.CustomerSubscription, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
// 	defer cancel()

// 	entity, err := s.repo.GetWithTx(ctx, dto.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	brokerRes := &customerBrokerContract.CustomerSubscription{
// 		ID:          entity.ID.String(),
// 		ServiceName: entity.ServiceName,
// 		Price:       entity.Price,
// 		UserID:      entity.UserID.String(),
// 		StartDate:   entity.StartDate.Format("01-2006"),
// 		EndDate:     conv.FormatPtrDate(entity.EndDate, "01-2006"),
// 	}
// 	// if entity.EndDate != nil {
// 	// 	brokerRes.EndDate = new(string)
// 	// 	*brokerRes.EndDate = entity.EndDate.Format("01-2006")
// 	// }
// 	return brokerRes, err
// }

// func (s *CustomerSubscriptionService) List(dto customerBrokerContract.CustomerSubscriptionsListRequest) (customerBrokerContract.CustomerSubscriptionsListResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
// 	defer cancel()

// 	entities, err := s.repo.ListWithTx(ctx, dto.UserID)

// 	var list []customerBrokerContract.CustomerSubscription
// 	for _, entity := range entities {
// 		list = append(list,
// 			customerBrokerContract.CustomerSubscription{
// 				ID:          entity.ID.String(),
// 				ServiceName: entity.ServiceName,
// 				Price:       entity.Price,
// 				UserID:      entity.UserID.String(),
// 				StartDate:   entity.StartDate.Format("01-2006"),
// 				EndDate:     conv.FormatPtrDate(entity.EndDate, "01-2006"),
// 			},
// 		)
// 	}

// 	return customerBrokerContract.CustomerSubscriptionsListResponse{
// 		Subscriptions: list,
// 	}, err
// }

// func (s *CustomerSubscriptionService) Delete(dto customerBrokerContract.CustomerSubscriptionDeleteRequest) (customerBrokerContract.CustomerSubscriptionDeleteResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
// 	defer cancel()

// 	err := s.repo.DeleteWithTx(ctx, dto.ID)
// 	return customerBrokerContract.CustomerSubscriptionDeleteResponse{}, err
// }

// func (s *CustomerSubscriptionService) Update(dto *customerBrokerContract.CustomerSubscriptionUpdateRequest) (customerBrokerContract.CustomerSubscriptionUpdateResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5000000*time.Millisecond)
// 	defer cancel()

// 	id, _ := uuid.Parse(dto.ID)

// 	// we do not retrieve entity from db previously
// 	// for performance reason
// 	entity := customerSubscriptionModel.NewCustomerSubscription(customerSubscriptionModel.NewCustomerSubscriptionDeps{
// 		ID: id,
// 	})
// 	mask := []string{}
// 	if dto.ServiceName != nil {
// 		entity.UpdateServiceName(dto.ServiceName)
// 		field := "ServiceName"
// 		mask = append(mask, field)
// 	}
// 	if dto.Price != nil {
// 		entity.UpdatePrice(dto.Price)
// 		field := "Price"
// 		mask = append(mask, field)
// 	}
// 	if dto.StartDate != nil {
// 		entity.UpdateStartDate(dto.StartDate)
// 		field := "StartDate"
// 		mask = append(mask, field)
// 	}
// 	if dto.EndDate != nil {
// 		entity.UpdateEndDate(dto.EndDate)
// 		field := "EndDate"
// 		mask = append(mask, field)
// 	}

// 	err := s.repo.UpdateWithTx(ctx, entity, mask)

// 	return customerBrokerContract.CustomerSubscriptionUpdateResponse{}, err
// }

// func (s *CustomerSubscriptionService) TotalCost(
// 	dto *customerBrokerContract.CustomerSubscriptionTotalCostRequest,
// ) (customerBrokerContract.CustomerSubscriptionTotalCostResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5000000*time.Millisecond)
// 	defer cancel()

// 	userID, _ := uuid.Parse(dto.UserID)
// 	entity := customerSubscriptionModel.NewCustomerSubscription(customerSubscriptionModel.NewCustomerSubscriptionDeps{
// 		UserID: userID,
// 	})
// 	mask := []string{}
// 	if dto.ServiceName != nil {
// 		entity.UpdateServiceName(dto.ServiceName)
// 		field := "ServiceName"
// 		mask = append(mask, field)
// 	}
// 	if dto.StartDate != nil {
// 		entity.UpdateStartDate(dto.StartDate)
// 		field := "StartDate"
// 		mask = append(mask, field)
// 	}
// 	if dto.EndDate != nil {
// 		entity.UpdateEndDate(dto.EndDate)
// 		field := "EndDate"
// 		mask = append(mask, field)
// 	}

// 	var total int
// 	total, err := s.repo.TotalCostWithTx(ctx, entity, mask)

// 	return customerBrokerContract.CustomerSubscriptionTotalCostResponse{Total: total}, err

// }

// // func (s *CustomerSubscriptionService) ListOfSubscriptions(userID string) ([]*customerSubscriptionModel.CustomerSubscription, error) {
// // 	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
// // 	defer cancel()

// // 	entities, err := s.repo.ListWithTx(ctx, userID)
// // 	return entities, err
// // }

// // func (s *CustomerSubscriptionService) AggregateSum(userID string) ([]*customerSubscriptionModel.CustomerSubscription, error) {
// // 	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Millisecond)
// // 	defer cancel()

// // 	entities, err := s.repo.AggregateSumWithTx(ctx, userID)
// // 	return entities, err
// // }
