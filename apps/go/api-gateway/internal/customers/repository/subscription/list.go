package customerSubscriptionsRepository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	model "project1.v0/api_gateway/internal/customers/models/subscription"
	customerDomainContract "project1.v0/contracts/domain/customer"
	customerBrokerContract "project1.v0/contracts/transport/broker/customer"
	"project1.v0/mappers/conv"
	db_postgresql "project1.v0/pkg/db/postgresql"
)

func (r *CustomerSubscriptionRepository) ListWithTx(
	ctx context.Context,
	dto *customerBrokerContract.CustomerSubscriptionsListRequest,
	fields []string,
) ([]*model.CustomerSubscription, error) {

	var entities []*model.CustomerSubscription

	businessErr := error(nil)

	err := db_postgresql.WithTx(r.db.DB, ctx, nil, func(tx *sql.Tx) error {
		var err error
		entities, err = r.List(ctx, tx, dto, fields)
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

func (r *CustomerSubscriptionRepository) List(
	ctx context.Context,
	tx *sql.Tx,
	dto *customerBrokerContract.CustomerSubscriptionsListRequest,
	fields []string,
) ([]*model.CustomerSubscription, error) {

	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`
	whereClauses := make([]string, 0, len(fields))
	args := make([]interface{}, 0, len(fields))

	mapNameToDb, err := r.decoder.GetDbFieldsStruct(&model.CustomerSubscription{})
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(dto).Elem()
	for _, field := range fields {
		if field == "Limit" || field == "Offset" {
			continue
		}
		dbFieldName, ok := mapNameToDb[field]
		if !ok {
			return nil, fmt.Errorf("field not found: %s", field)
		}
		fieldVal := v.FieldByName(field)
		if !fieldVal.IsValid() || fieldVal.IsNil() {
			continue
		}
		whereClauses = append(whereClauses, fmt.Sprintf("%s=$%d", dbFieldName, len(args)+1))

		val := fieldVal.Elem().Interface()

		idx := len(args) + 1
		switch field {
		case "ServiceName":
			whereClauses = append(whereClauses, fmt.Sprintf("%s ILIKE $%d", dbFieldName, idx))
			strVal, _ := val.(string)
			args = append(args, fmt.Sprintf("%%%s%%", strVal))
		case "StartDate":
			whereClauses = append(whereClauses, fmt.Sprintf("%s >= $%d", dbFieldName, idx))
			strVal, _ := val.(string)
			args = append(args, conv.ParseDate(strVal))
		case "EndDate":
			whereClauses = append(whereClauses, fmt.Sprintf("%s <= $%d", dbFieldName, idx))
			strVal, _ := val.(string)
			args = append(args, conv.ParseDate(strVal))
		default:
			whereClauses = append(whereClauses, fmt.Sprintf("%s=$%d", dbFieldName, idx))
			args = append(args, val)
		}
		args = append(args, val)
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if dto.Limit != nil {
		args = append(args, *dto.Limit)
		query += fmt.Sprintf(" LIMIT $%d", len(args))
	}
	if dto.Offset != nil {
		args = append(args, *dto.Offset)
		query += fmt.Sprintf(" OFFSET $%d", len(args))
	}

	rows, err := tx.QueryContext(ctx, query, args...)
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
