package customerSubscriptionModel

import (
	"time"

	"github.com/google/uuid"
	conv "project1.v0/mappers/conv"
)

type CustomerSubscription struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	ServiceName string     `json:"service_name" db:"service_name"`
	Price       int        `json:"price" db:"price"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"`
	StartDate   time.Time  `json:"start_date" db:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty" db:"end_date"`
}

func (s *CustomerSubscription) TableName() string {
	return "subscription"
}

type NewCustomerSubscriptionDeps struct {
	ID          uuid.UUID
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}

func NewCustomerSubscription(deps NewCustomerSubscriptionDeps) *CustomerSubscription {
	return &CustomerSubscription{
		ID:          deps.ID,
		ServiceName: deps.ServiceName,
		Price:       deps.Price,
		UserID:      deps.UserID,
		StartDate:   deps.StartDate,
		EndDate:     deps.EndDate,
	}
}

func (s *CustomerSubscription) UpdateUserID(userIDStr *string) error {
	userID, err := uuid.Parse(*userIDStr)
	if err != nil {
		return err
	}
	s.UserID = userID
	return nil
}

func (s *CustomerSubscription) UpdateServiceName(ServiceName *string) {
	s.ServiceName = *ServiceName
}
func (s *CustomerSubscription) UpdatePrice(Price *int) {
	s.Price = *Price
}
func (s *CustomerSubscription) UpdateStartDate(startDate *string) {
	s.StartDate = *conv.ParseDatePtr(startDate)
}
func (s *CustomerSubscription) UpdateEndDate(endDate *string) {
	s.EndDate = conv.ParseDatePtr(endDate)
}
