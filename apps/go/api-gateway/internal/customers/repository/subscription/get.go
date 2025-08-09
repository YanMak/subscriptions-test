package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"errors"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	customerDomainContract "project1.v0/contracts/domain/customer"

	//customerDomainContract "project1.v0/contracts/domain/customer"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) GetWithTx(ctx context.Context, id string) (*model.CustomerSubscription, error) {

	var entity *model.CustomerSubscription
	businessErr := error(nil)

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		var err error
		entity, err = r.Get(ctx, tx, id)
		if errors.Is(err, sql.ErrNoRows) {
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
	return entity, nil
}

func (r *CustomerSubscriptionRepository) Get(ctx context.Context, tx *sql.Tx, id string) (*model.CustomerSubscription, error) {
	var s model.CustomerSubscription
	err := tx.QueryRowContext(ctx,
		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id = $1`, id,
	).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	return &s, err
}

// func (r *CustomerSubscriptionRepository) dbFieldName(field string) (string, error) {
// 	for _, fi := range r.decoder.Fields {
// 		if fi.Name == field {
// 			if dbTag, ok := fi.TagMap["db"]; ok {
// 				return dbTag, nil
// 			}
// 			break
// 		}
// 	}
// 	return "", fmt.Errorf("field not found: %s", field)
// }
