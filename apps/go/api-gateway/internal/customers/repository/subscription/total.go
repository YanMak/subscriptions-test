package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) TotalCostWithTx(ctx context.Context, s *model.CustomerSubscription, fields []string) (int, error) {

	var total int
	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		var err error
		total, err = r.TotalCost(ctx, tx, s, fields)
		if err != nil {
			return err // rollback
		}
		return nil // commit
	})
	if err != nil {
		return 0, err
	}

	return total, err
}

// func (r *CustomerSubscriptionRepository) AggregateSum(ctx context.Context, tx *sql.Tx, userID, serviceName string, from, to time.Time) (int, error) {
// 	var sum int
// 	err := tx.QueryRowContext(ctx,
// 		`SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE user_id=$1 AND service_name=$2 AND start_date >= $3 AND (end_date<= $4 OR end_date IS NULL)`,
// 		userID, serviceName, from, to,
// 	).Scan(&sum)
// 	return sum, err
// }

func (r *CustomerSubscriptionRepository) TotalCost(ctx context.Context, tx *sql.Tx, entity *model.CustomerSubscription, fields []string) (int, error) {
	var totalCost int
	var err error

	query := "SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	val := reflect.ValueOf(entity).Elem()
	var idx int

	mapNameToDb, err := r.decoder.GetDbFieldsStruct(entity)
	if err != nil {
		return 0, err
	}

	for _, field := range fields {
		dbFieldName, ok := mapNameToDb[field]
		if !ok {
			return 0, fmt.Errorf("field not found: %s", field)
		}

		fieldVal := val.FieldByName(field)
		if !fieldVal.IsValid() {
			return 0, fmt.Errorf("field not found: %s", field)
		}

		switch {
		case field == "ServiceName":
			idx++
			query += fmt.Sprintf(" AND %s ILIKE $%d", dbFieldName, idx)
			args = append(args, fmt.Sprintf("%%%s%%", fieldVal.Interface()))

		case field == "StartDate":
			idx++
			query += fmt.Sprintf(" AND %s >= $%d", dbFieldName, idx)
			args = append(args, fieldVal.Interface())

		case field == "EndDate":
			idx++
			query += fmt.Sprintf(" AND %s <= $%d", dbFieldName, idx)
			args = append(args, fieldVal.Interface())
		default:
			idx++
			query += fmt.Sprintf(" AND %s = $%d", dbFieldName, idx)
			args = append(args, fieldVal.Interface())
		}

	}

	err = tx.QueryRowContext(ctx, query, args...).Scan(&totalCost)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total cost: %w", err)
	}

	return totalCost, nil

	// // Фильтрация по user_id, если задан
	// if entity.UserID != nil {
	// 	query += " AND user_id = $1"
	// 	args = append(args, *userID)
	// }

	// // Фильтрация по service_name, если задан
	// if entity.ServiceName != nil {
	// 	query += " AND service_name ILIKE $2"
	// 	args = append(args, "%"+*serviceName+"%")
	// }

	// // Фильтрация по start_date, если задан
	// if entity.StartDate != nil {
	// 	query += " AND start_date >= $3"
	// 	args = append(args, entity.StartDate)
	// }

	// // Фильтрация по end_date, если задан
	// if entity.EndDate != nil {
	// 	query += " AND end_date <= $4"
	// 	args = append(args, entity.EndDate)
	// }

}
