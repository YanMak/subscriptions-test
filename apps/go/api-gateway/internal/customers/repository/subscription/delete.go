package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"errors"

	customerDomainContract "project1.v0/contracts/domain/customer"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) DeleteWithTx(ctx context.Context, id string) error {

	businessErr := error(nil)

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		err := r.Delete(ctx, tx, id)
		if errors.Is(err, sql.ErrNoRows) {
			businessErr = customerDomainContract.ErrSubscriptionNotFound
			return nil
		}
		return err
	})
	if businessErr != nil {
		return businessErr
	}
	if err != nil {
		return err
	}

	return nil

}

func (r *CustomerSubscriptionRepository) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	_, err := tx.ExecContext(ctx,
		`DELETE FROM subscriptions WHERE id=$1`, id,
	)
	return err
}
