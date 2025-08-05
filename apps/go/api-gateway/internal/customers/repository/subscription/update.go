package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) UpdateWithTx(ctx context.Context, s *model.CustomerSubscription, fields []string) error {

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		//return r.Update(ctx, tx, s, fields)
		err := r.Update(ctx, tx, s, fields)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerSubscriptionRepository) Update(ctx context.Context, tx *sql.Tx, s *model.CustomerSubscription, fields []string) error {

	if len(fields) == 0 {
		return nil
	}

	setClauses := make([]string, 0, len(fields))
	args := make([]interface{}, 0, len(fields)+1)

	mapNameToDb, err := r.decoder.GetDbFieldsStruct(s)
	if err != nil {
		return err
	}
	val := reflect.ValueOf(s).Elem()

	for idx, field := range fields {
		dbFieldName, ok := mapNameToDb[field]
		if !ok {
			return fmt.Errorf("field not found: %s", field)
		}
		setClauses = append(setClauses, fmt.Sprintf("%s=$%d", dbFieldName, idx+1))

		fieldVal := val.FieldByName(field)
		if !fieldVal.IsValid() {
			return fmt.Errorf("field not found: %s", field)
		}
		args = append(args, fieldVal.Interface())
	}

	// val := reflect.ValueOf(s).Elem()

	// for idx, field := range fields {
	// 	dbFieldName := r.DtosToDbMapService.GetFieldMapping(tableName, field)
	// 	setClauses = append(setClauses, fmt.Sprintf("%s=$%d", dbFieldName, idx+1))

	// 	fieldVal := val.FieldByName(field)
	// 	if !fieldVal.IsValid() {
	// 		return fmt.Errorf("field not found: %s", field)
	// 	}
	// 	args = append(args, fieldVal.Interface())
	// }

	query := fmt.Sprintf(
		`UPDATE subscriptions SET %s WHERE id=$%d`,
		strings.Join(setClauses, ", "),
		len(fields)+1,
	)
	args = append(args, s.ID)

	_, err = tx.ExecContext(ctx, query, args...)
	return err

}

// func (r *CustomerSubscriptionRepository) UpdateArchive(ctx context.Context, tx *sql.Tx, s *model.CustomerSubscription, fields []string) error {

// 	if len(fields) == 0 {
// 		return nil
// 	}

// 	setClauses := make([]string, 0, len(fields))
// 	args := make([]interface{}, 0, len(fields)+1)
// 	val := reflect.ValueOf(s).Elem()
// 	tableName := s.TableName()

// 	for idx, field := range fields {
// 		dbFieldName := r.DtosToDbMapService.GetFieldMapping(tableName, field)
// 		setClauses = append(setClauses, fmt.Sprintf("%s=$%d", dbFieldName, idx+1))

// 		fieldVal := val.FieldByName(field)
// 		if !fieldVal.IsValid() {
// 			return fmt.Errorf("field not found: %s", field)
// 		}
// 		args = append(args, fieldVal.Interface())
// 	}

// 	query := fmt.Sprintf(
// 		`UPDATE subscriptions SET %s WHERE id=$%d`,
// 		strings.Join(setClauses, ", "),
// 		len(fields)+1,
// 	)
// 	args = append(args, s.ID)

// 	_, err := tx.ExecContext(ctx, query, args...)
// 	return err

// }
