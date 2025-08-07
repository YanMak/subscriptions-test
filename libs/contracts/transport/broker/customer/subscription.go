package customerBrokerContract

import uuid "github.com/google/uuid"

// /////////////
// CREATE Subscription
const (
	SubscriptionCreateTopic = "customer.create-subscription.command"
)

type CustomerSubscriptionCreateRequest struct {
	ServiceName string  `json:"service_name" validate:"required"`                   // @example "Yandex Plus at broker"
	Price       int     `json:"price" validate:"required"`                          // @example 400
	UserID      string  `json:"user_id" validate:"required,uuid"`                   // @example "60601fee-2bf1-4721-ae6f-7636e79a0cba"
	StartDate   string  `json:"start_date" validate:"required,date_format_mm_yyyy"` // @example "07-2025"
	EndDate     *string `json:"end_date,omitempty"`                                 // @example "07-2026"
}

type CustomerSubscriptionCreateResponse struct {
	ID string `json:"id"`
}

// /////////////
// Update Subscription
const (
	SubscriptionUpdateTopic = "customer.update-subscription.command"
)

type CustomerSubscriptionUpdateRequest struct {
	ID          string  `json:"id" validate:"required"`
	UserID      string  `json:"user_id" validate:"required,uuid"`
	ServiceName *string `json:"service_name,omitempty"`
	Price       *int    `json:"price,omitempty" validate:"omitempty"`
	StartDate   *string `json:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	EndDate     *string `json:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
}

type CustomerSubscriptionUpdateResponse struct{}

// /////////////
// Get Subscription

type CustomerSubscriptionIDRequest struct {
	ID string `json:"id"`
}

const (
	SubscriptionGetTopic = "customer.get-subscription.query"
)

// CustomerSubscriptionIDRequest as request

type CustomerSubscription struct {
	ID          string  `json:"id" db:"id"`
	ServiceName string  `json:"service_name" db:"service_name"`
	Price       int     `json:"price" db:"price"`
	UserID      string  `json:"user_id" db:"user_id"`
	StartDate   string  `json:"start_date" db:"start_date"`
	EndDate     *string `json:"end_date,omitempty" db:"end_date"`
}

// /////////////
// Get Subscriptions list
const (
	SubscriptionsListTopic = "customer.get-subscriptions-list.query"
)

type CustomerSubscriptionsListRequest struct {
	UserID      *string      `json:"user_id,omitempty" validate:"omitempty,uuid"`
	UserIDs     *[]uuid.UUID `json:"user_ids,omitempty" validate:"omitempty,dive,uuid"`
	ServiceName *string      `json:"service_name,omitempty"`
	StartDate   *string      `json:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	EndDate     *string      `json:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	PriceFrom   *int         `json:"price_from,omitempty" validate:"omitempty,min=0"`
	PriceTo     *int         `json:"price_to,omitempty" validate:"omitempty,min=0"`

	Limit  *int `json:"limit,omitempty" validate:"omitempty"`
	Offset *int `json:"offset,omitempty" validate:"omitempty"`

	Sort []struct {
		Field     string `json:"field"`
		Direction string `json:"direction"`
	} `json:"sort,omitempty"`
}

type CustomerSubscriptionsListResponse struct {
	Subscriptions []CustomerSubscription `json:"subscriptions"`
}

// /////////////
// Delete Subscription
const (
	SubscriptionDeleteTopic = "customer.delete-subscription.command"
)

type CustomerSubscriptionDeleteResponse struct{}

// /////////////
// TotalCost Subscription
const (
	SubscriptionTotalCostTopic = "customer.total-cost-subscription.query"
)

// CustomerSubscriptionIDRequest as request

type CustomerSubscriptionTotalCostRequest struct {
	UserID      *string `json:"user_id,omitempty" validate:"required,uuid"`
	ServiceName *string `json:"service_name,omitempty"`
	StartDate   *string `json:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	EndDate     *string `json:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	PriceFrom   *int    `json:"price_from,omitempty" validate:"omitempty,min=0"`
	PriceTo     *int    `json:"price_to,omitempty" validate:"omitempty,min=0"`
}

type CustomerSubscriptionTotalCostResponse struct {
	Total int `json:"total"`
}
