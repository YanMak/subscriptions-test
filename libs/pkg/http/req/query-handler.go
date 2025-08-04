package req

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	validatorХ "project1.v0/pkg/validator_x"
)

func HandleQuery[T any](r *http.Request, vld *validatorХ.ValidatorX) (*T, error) {
	var out T
	values := r.URL.Query()

	err := mapQueryToStruct(values, &out)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	if err := vld.Validate.Struct(out); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return &out, nil
}

func mapQueryToStruct(values url.Values, dst any) error {
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		paramName := strings.Split(tag, ",")[0]

		queryVal := values.Get(paramName)
		if queryVal == "" {
			continue
		}

		fv := v.Field(i)
		switch fv.Kind() {
		case reflect.Ptr:
			switch fv.Type().Elem().Kind() {
			case reflect.String:
				fv.Set(reflect.ValueOf(&queryVal))
			case reflect.Int:
				i, err := strconv.Atoi(queryVal)
				if err != nil {
					return fmt.Errorf("invalid int for field %s", field.Name)
				}
				fv.Set(reflect.ValueOf(&i))
			}
		case reflect.String:
			fv.SetString(queryVal)
		case reflect.Int:
			i, err := strconv.Atoi(queryVal)
			if err != nil {
				return fmt.Errorf("invalid int for field %s", field.Name)
			}
			fv.SetInt(int64(i))
		}
	}
	return nil
}
