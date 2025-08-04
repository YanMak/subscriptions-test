package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	customerDomainContract "project1.v0/contracts/domain/customer"

	//customerDomainContract "project1.v0/contracts/domain/customer"
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

// func (r *CustomerSubscriptionRepository) GetWithTx(ctx context.Context, id string) (*model.CustomerSubscription, error) {

// 	var entity *model.CustomerSubscription
// 	businessErr := error(nil)

// 	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
// 		var err error
// 		entity, err = r.Get(ctx, tx, id)
// 		if errors.Is(err, sql.ErrNoRows) {
// 			businessErr = customerDomainContract.ErrSubscriptionNotFound
// 			return nil
// 		}
// 		return err
// 	})
// 	if businessErr != nil {
// 		return nil, businessErr
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

func (r *CustomerSubscriptionRepository) UpdateWithTx(ctx context.Context, s *model.CustomerSubscription, fields []string) error {

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		//return r.Update(ctx, tx, s, fields)
		err := r.Update(ctx, tx, s, fields)
		return err
	})
	if err != nil {
		// error processing on rollback
	}

	return nil
}

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

// func (r *CustomerSubscriptionRepository) ListWithTx(ctx context.Context, userID string) ([]*model.CustomerSubscription, error) {

// 	var entities []*model.CustomerSubscription

// 	businessErr := error(nil)

// 	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
// 		var err error
// 		entities, err = r.List(ctx, tx, userID)
// 		if len(entities) == 0 && err == nil {
// 			businessErr = customerDomainContract.ErrSubscriptionNotFound
// 			return nil
// 		}
// 		return err
// 	})
// 	if businessErr != nil {
// 		return nil, businessErr
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return entities, err
// }

// func (r *CustomerSubscriptionRepository) TotalCostWithTx(ctx context.Context, s *model.CustomerSubscription, fields []string) (int, error) {

// 	var total int
// 	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
// 		var err error
// 		total, err = r.TotalCost(ctx, tx, s, fields)
// 		if err != nil {
// 			return err // rollback
// 		}
// 		return nil // commit
// 	})
// 	if err != nil {
// 		// error processing on rollback
// 	}

// 	return total, err
// }

func (r *CustomerSubscriptionRepository) Create(ctx context.Context, tx *sql.Tx, s *model.CustomerSubscription) (string, error) {
	var id string
	err := tx.QueryRowContext(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate,
	).Scan(&id)
	return id, err
}

// func (r *CustomerSubscriptionRepository) Get(ctx context.Context, tx *sql.Tx, id string) (*model.CustomerSubscription, error) {
// 	var s model.CustomerSubscription
// 	err := tx.QueryRowContext(ctx,
// 		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id = $1`, id,
// 	).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
// 	return &s, err
// }

func (r *CustomerSubscriptionRepository) Update(ctx context.Context, tx *sql.Tx, s *model.CustomerSubscription, fields []string) error {

	if len(fields) == 0 {
		return nil
	}

	setClauses := make([]string, 0, len(fields))
	args := make([]interface{}, 0, len(fields)+1)
	val := reflect.ValueOf(s).Elem()
	tableName := s.TableName()

	for idx, field := range fields {
		dbFieldName := r.DtosToDbMapService.GetFieldMapping(tableName, field)
		setClauses = append(setClauses, fmt.Sprintf("%s=$%d", dbFieldName, idx+1))

		fieldVal := val.FieldByName(field)
		if !fieldVal.IsValid() {
			return fmt.Errorf("field not found: %s", field)
		}
		args = append(args, fieldVal.Interface())
	}

	query := fmt.Sprintf(
		`UPDATE subscriptions SET %s WHERE id=$%d`,
		strings.Join(setClauses, ", "),
		len(fields)+1,
	)
	args = append(args, s.ID)

	_, err := tx.ExecContext(ctx, query, args...)
	return err

}

func (r *CustomerSubscriptionRepository) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	_, err := tx.ExecContext(ctx,
		`DELETE FROM subscriptions WHERE id=$1`, id,
	)
	return err
}

// func (r *CustomerSubscriptionRepository) List(ctx context.Context, tx *sql.Tx, userID string) ([]*model.CustomerSubscription, int, error) {
// 	rows, err := tx.QueryContext(ctx, `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE user_id=$1`, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var subs []*model.CustomerSubscription
// 	for rows.Next() {
// 		var s model.CustomerSubscription
// 		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
// 			return nil, err
// 		}
// 		subs = append(subs, &s)
// 	}
// 	return subs, nil
// }

// func (r *CustomerSubscriptionRepository) AggregateSum(ctx context.Context, tx *sql.Tx, userID, serviceName string, from, to time.Time) (int, error) {
// 	var sum int
// 	err := tx.QueryRowContext(ctx,
// 		`SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE user_id=$1 AND service_name=$2 AND start_date >= $3 AND (end_date<= $4 OR end_date IS NULL)`,
// 		userID, serviceName, from, to,
// 	).Scan(&sum)
// 	return sum, err
// }

// func (r *CustomerSubscriptionRepository) TotalCost(ctx context.Context, tx *sql.Tx, entity *model.CustomerSubscription, fields []string) (int, error) {
// 	var totalCost int
// 	var err error

// 	query := "SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE 1=1"
// 	args := []interface{}{}
// 	val := reflect.ValueOf(entity).Elem()
// 	tableName := entity.TableName()
// 	var idx int

// 	// implement UserID
// 	idx++
// 	field := "UserID"
// 	fieldVal := val.FieldByName(field)
// 	UserIdDbFieldName := r.DtosToDbMapService.GetFieldMapping(tableName, field)
// 	query += fmt.Sprintf(" AND %s = $%d", UserIdDbFieldName, idx)
// 	args = append(args, fieldVal.Interface())

// 	for _, field := range fields {
// 		dbFieldName := r.DtosToDbMapService.GetFieldMapping(tableName, field)
// 		fieldVal := val.FieldByName(field)
// 		if !fieldVal.IsValid() {
// 			return 0, fmt.Errorf("field not found: %s", field)
// 		}

// 		switch {
// 		case field == "ServiceName":
// 			idx++
// 			query += fmt.Sprintf(" AND %s ILIKE $%d", dbFieldName, idx)
// 			args = append(args, fmt.Sprintf("%%%s%%", fieldVal.Interface()))

// 		case field == "StartDate":
// 			idx++
// 			query += fmt.Sprintf(" AND %s >= $%d", dbFieldName, idx)
// 			args = append(args, fieldVal.Interface())

// 		case field == "EndDate":
// 			idx++
// 			query += fmt.Sprintf(" AND %s <= $%d", dbFieldName, idx)
// 			args = append(args, fieldVal.Interface())
// 		}

// 	}

// 	err = tx.QueryRowContext(ctx, query, args...).Scan(&totalCost)
// 	if err != nil {
// 		return 0, fmt.Errorf("failed to calculate total cost: %w", err)
// 	}

// 	return totalCost, nil

// 	// // Фильтрация по user_id, если задан
// 	// if entity.UserID != nil {
// 	// 	query += " AND user_id = $1"
// 	// 	args = append(args, *userID)
// 	// }

// 	// // Фильтрация по service_name, если задан
// 	// if entity.ServiceName != nil {
// 	// 	query += " AND service_name ILIKE $2"
// 	// 	args = append(args, "%"+*serviceName+"%")
// 	// }

// 	// // Фильтрация по start_date, если задан
// 	// if entity.StartDate != nil {
// 	// 	query += " AND start_date >= $3"
// 	// 	args = append(args, entity.StartDate)
// 	// }

// 	// // Фильтрация по end_date, если задан
// 	// if entity.EndDate != nil {
// 	// 	query += " AND end_date <= $4"
// 	// 	args = append(args, entity.EndDate)
// 	// }

// }
