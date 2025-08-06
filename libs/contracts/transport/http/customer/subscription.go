package customerHttpContract

// ///////
// CREATE

type CustomerSubscriptionCreateRequest struct {
	ServiceName string  `json:"service_name" validate:"required"`                            // @example "Yandex Plus"
	Price       int     `json:"price" validate:"required"`                                   // @example 400
	UserID      string  `json:"user_id" validate:"required,uuid"`                            // @example "60601fee-2bf1-4721-ae6f-7636e79a0cba"
	StartDate   string  `json:"start_date" validate:"required,date_format_mm_yyyy"`          // @example "07-2025"
	EndDate     *string `json:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"` // @example "07-2026"
}

type CustomerSubscriptionCreateResponse struct {
	ID string `json:"id"`
}

// Get, List
type CustomerSubscriptionResponse struct {
	ID          string  `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}

/////////
// Update

type CustomerSubscriptionUpdateRequest struct {
	ServiceName *string `json:"service_name,omitempty"`                                        // @example "Yandex Plus"
	Price       *int    `json:"price,omitempty"`                                               // @example 400
	StartDate   *string `json:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"` // @example "07-2025"
	EndDate     *string `json:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`   // @example "07-2026"
}
type CustomerSubscriptionUpdatePath struct {
	ID *string `path:"id" validate:"uuid"` //
}

// type CustomerSubscriptionUpdateQuery struct {
// 	UserID    *string `query:"user_id" validate:"uuid"`                                       // @example "60601fee-2bf1-4721-ae6f-7636e79a0cba"
// 	StartDate *string `query:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"` // @example "07-2025"
// }

// Response is empty

// ///////
// List
type SubscriptionListQuery struct {
	UserID      *string `query:"user_id,omitempty" validate:"omitempty,uuid"`
	ServiceName *string `query:"service_name,omitempty" validate:"omitempty"`
	StartDate   *string `query:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	EndDate     *string `query:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	PriceFrom   *int    `query:"price_from,omitempty" validate:"omitempty,min=0"`
	PriceTo     *int    `query:"price_to,omitempty" validate:"omitempty,min=0"`
	Limit       *int    `query:"limit,omitempty" validate:"omitempty,min=1,max=100"`
	Offset      *int    `query:"offset,omitempty" validate:"omitempty,min=0"`
}

/////////
// Aggregate

// @Description Aggregate request: subscriptions sum
// @name SubscriptionAggregateRequest
type SubscriptionTotalQuery struct {
	UserID      *string `query:"user_id" validate:"required,uuid"`
	ServiceName *string `query:"service_name,omitempty" validate:"omitempty"`
	StartDate   *string `query:"start_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	EndDate     *string `query:"end_date,omitempty" validate:"omitempty,date_format_mm_yyyy"`
	PriceFrom   *int    `query:"price_from,omitempty" validate:"omitempty,min=0"`
	PriceTo     *int    `query:"price_to,omitempty" validate:"omitempty,min=0"`
}

// @Description Aggregate response: сумма
// @name SubscriptionAggregateResponse
type SubscriptionTotalResponse struct {
	Total int `json:"total"` // Суммарная стоимость подписок (за период, по фильтрам)
}

type CustomerSubscriptionIDPath struct {
	ID *string `path:"id" validate:"uuid"` //
}

// ErrorResponse — commonly used struct for err
// @Description error response body
// @name ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
