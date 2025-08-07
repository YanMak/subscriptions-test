package customerDomainContract

import uuid "github.com/google/uuid"

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

	SortBy    *string `json:"sort_by,omitempty"`
	SortOrder *string `json:"sort_order,omitempty"`
}
