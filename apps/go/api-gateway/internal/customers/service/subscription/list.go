package customerSubscriptionService

import (
	"context"
	"time"

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
			Field     string `json:"field" validate:"required,oneof=service_name start_date end_date price"`
			Direction string `json:"direction" validate:"required,oneof=asc desc"`
		}{
			Field:     string(customerDomainContract.SubscriptionSortField(srt.Field)),
			Direction: string(srt.Direction), // or string(domainContract.SortDirection(srt.Direction)) if SortDirection is defined
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
