package customerSubscriptionService

import (
	"context"
	"time"

	domainContract "project1.v0/contracts/domain"
	customerDomainContract "project1.v0/contracts/domain/customer"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	conv "project1.v0/mappers/conv"
)

func (s *CustomerSubscriptionService) List(dto *customerBrokerContract.CustomerSubscriptionsListRequest) (customerBrokerContract.CustomerSubscriptionsListResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500000*time.Millisecond)
	defer cancel()

	domainDto := &customerDomainContract.CustomerSubscriptionsListRequest{
		UserID:      dto.UserID,
		UserIDs:     dto.UserIDs,
		ServiceName: dto.ServiceName,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		PriceFrom:   dto.PriceFrom,
		PriceTo:     dto.PriceTo,
		Limit:       dto.Limit,
		Offset:      dto.Offset,
	}

	for _, srt := range dto.Sort {
		domainDto.Sort = append(domainDto.Sort, struct {
			Field     customerDomainContract.SubscriptionSortField `json:"field" validate:"required,oneof=service_name start_date end_date price"`
			Direction domainContract.SortDirection                 `json:"direction" validate:"required,oneof=asc desc"`
		}{
			Field:     customerDomainContract.SubscriptionSortField(srt.Field),
			Direction: domainContract.SortDirection(srt.Direction),
		})
	}

	entities, err := s.repo.ListWithTx(ctx, domainDto)

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
