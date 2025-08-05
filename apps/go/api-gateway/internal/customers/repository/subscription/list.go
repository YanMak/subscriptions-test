package customerSubscriptionsRepository

import (
	"context"
	"database/sql"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	customerDomainContract "project1.v0/contracts/domain/customer"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) ListWithTx(ctx context.Context, userID string) ([]*model.CustomerSubscription, error) {

	var entities []*model.CustomerSubscription

	businessErr := error(nil)

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		var err error
		entities, err = r.List(ctx, tx, userID)
		if len(entities) == 0 && err == nil {
			businessErr = customerDomainContract.ErrSubscriptionNotFound
			return nil
		}
		return err
	})
	if businessErr != nil {
		return nil, businessErr
	}
	if err != nil {
		return nil, err
	}

	return entities, err
}

func (r *CustomerSubscriptionRepository) List(ctx context.Context, tx *sql.Tx, userID string) ([]*model.CustomerSubscription, error) {
	rows, err := tx.QueryContext(ctx, `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subs []*model.CustomerSubscription
	for rows.Next() {
		var s model.CustomerSubscription
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			return nil, err
		}
		subs = append(subs, &s)
	}
	return subs, nil
}
