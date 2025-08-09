package customerSubscriptionsRepository

import (
	"context"
	"database/sql"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) CreateWithTx(ctx context.Context, s *model.CustomerSubscription) (string, error) {

	var subscriptionID string

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		var err error
		subscriptionID, err = r.Create(ctx, tx, &model.CustomerSubscription{
			ServiceName: s.ServiceName,
			Price:       s.Price,
			UserID:      s.UserID,
			StartDate:   s.StartDate,
			EndDate:     s.EndDate,
		})
		if err != nil {
			return err // rollback
		}
		return nil // commit
	})
	if err != nil {
		// error processing on rollback
		return "", err
	}

	return subscriptionID, nil
}

func (r *CustomerSubscriptionRepository) Create(ctx context.Context, tx *sql.Tx, s *model.CustomerSubscription) (string, error) {
	var id string
	err := tx.QueryRowContext(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate,
	).Scan(&id)
	return id, err
}
