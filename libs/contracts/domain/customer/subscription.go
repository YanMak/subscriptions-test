package customerDomainContract

import (
	uuid "github.com/google/uuid"
)

type SubscriptionSortField string

const (
	SubscriptionSortFieldServiceName SubscriptionSortField = "service_name"
	SubscriptionSortFieldStartDate   SubscriptionSortField = "start_date"
	SubscriptionSortFieldEndDate     SubscriptionSortField = "end_date"
	SubscriptionSortFieldPrice       SubscriptionSortField = "price"
)

// type Sort []struct {
// 	Field     SubscriptionSortField `json:"field"`
// 	Direction domain.SortDirection  `json:"direction"`
// }

type CustomerSubscriptionIDRequest struct {
	ID string `json:"id" db:"id"`
}

type CustomerSubscriptionsListRequest struct {
	UserID      *string      `db:"user_id"`
	UserIDs     *[]uuid.UUID `db:"user_id"`
	ServiceName *string      `db:"service_name"`
	StartDate   *string      `db:"start_date"`
	EndDate     *string      `db:"end_date"`
	PriceFrom   *int         `db:"price"`
	PriceTo     *int         `db:"price"`

	Limit  *int
	Offset *int

	Sort []struct {
		Field     string `json:"field" validate:"required,oneof=service_name start_date end_date price"`
		Direction string `json:"direction" validate:"required,oneof=asc desc"`
	} `json:"sort,omitempty" validate:"omitempty,dive"`
}
