package middleware

// // Body
// func decodeBody[T any](r *http.Request, out *T) error {
// 	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
// 		return fmt.Errorf("decode body: %w", err)
// 	}
// 	return nil
// }

// // Query params
// func mapQueryToStruct[T any](r *http.Request, out *T) error {
// 	values := r.URL.Query()

// 	v := reflect.ValueOf(out)
// 	if v.Kind() != reflect.Ptr || v.IsNil() {
// 		return errors.New("mapQueryToStruct: out must be a non-nil pointer")
// 	}
// 	v = v.Elem()

// 	t := v.Type()
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)

// 		// Получаем имя параметра из тега
// 		paramName := field.Tag.Get("query")
// 		if paramName == "" {
// 			paramName = field.Tag.Get("json")
// 		}
// 		if paramName == "" {
// 			paramName = strings.ToLower(field.Name)
// 		}
// 		paramName = strings.Split(paramName, ",")[0] // убрать `omitempty`

// 		rawVal := values.Get(paramName)
// 		if rawVal == "" {
// 			continue
// 		}

// 		fieldVal := v.Field(i)
// 		if !fieldVal.CanSet() {
// 			continue
// 		}

// 		switch field.Type.Kind() {
// 		case reflect.String:
// 			fieldVal.SetString(rawVal)
// 		case reflect.Int:
// 			if intVal, err := strconv.Atoi(rawVal); err == nil {
// 				fieldVal.SetInt(int64(intVal))
// 			} else {
// 				return fmt.Errorf("invalid int for query param %s", paramName)
// 			}
// 		case reflect.Ptr:
// 			switch field.Type.Elem().Kind() {
// 			case reflect.String:
// 				fieldVal.Set(reflect.ValueOf(&rawVal))
// 			case reflect.Int:
// 				if intVal, err := strconv.Atoi(rawVal); err == nil {
// 					fieldVal.Set(reflect.ValueOf(&intVal))
// 				} else {
// 					return fmt.Errorf("invalid *int for query param %s", paramName)
// 				}
// 			}
// 		default:
// 			return fmt.Errorf("unsupported type %s for query field %s", field.Type.Kind(), field.Name)
// 		}
// 	}

// 	return nil
// }

// // Path
// func mapPathParamsToStruct[T any](r *http.Request, out *T) error {

// 	v := reflect.ValueOf(out)
// 	if v.Kind() != reflect.Ptr || v.IsNil() {
// 		return errors.New("mapPathParamsToStruct: out must be a non-nil pointer")
// 	}
// 	v = v.Elem()

// 	t := v.Type()
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		tag := field.Tag.Get("path")
// 		var paramName string
// 		if tag == "" || tag == "-" {
// 			paramName = strings.ToLower(field.Name)
// 		} else {
// 			paramName = strings.Split(tag, ",")[0]
// 		}

// 		rawVal := r.PathValue(paramName)
// 		if rawVal == "" {
// 			continue
// 		}

// 		fieldVal := v.Field(i)
// 		if !fieldVal.CanSet() {
// 			continue
// 		}

// 		switch field.Type.Kind() {
// 		case reflect.String:
// 			fieldVal.SetString(rawVal)
// 		case reflect.Int:
// 			if intVal, err := strconv.Atoi(rawVal); err == nil {
// 				fieldVal.SetInt(int64(intVal))
// 			} else {
// 				return fmt.Errorf("invalid int for path param %s", paramName)
// 			}
// 		case reflect.Ptr:
// 			// support for *string, *int
// 			switch field.Type.Elem().Kind() {
// 			case reflect.String:
// 				fieldVal.Set(reflect.ValueOf(&rawVal))
// 			case reflect.Int:
// 				if intVal, err := strconv.Atoi(rawVal); err == nil {
// 					fieldVal.Set(reflect.ValueOf(&intVal))
// 				} else {
// 					return fmt.Errorf("invalid *int for path param %s", paramName)
// 				}
// 			}
// 		default:
// 			return fmt.Errorf("unsupported type %s for field %s", field.Type.Kind(), field.Name)
// 		}
// 	}

// 	return nil
// }
