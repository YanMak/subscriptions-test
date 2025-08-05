package decoder

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

func (m *DecoderService) MapQueryToStruct(values url.Values, out any) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("output must be a non-nil pointer")
	}
	v = v.Elem()
	reflectedType := v.Type()
	meta := m.GetStructMeta(reflectedType)

	for tagVal, fi := range meta.FieldsByTag["query"] {
		queryVal := values.Get(tagVal)
		if queryVal == "" {
			continue
		}
		field := v.Field(fi.Index)
		if err := fi.SetFunc(field, queryVal); err != nil {
			return fmt.Errorf("failed to set field %s: %w", fi.Name, err)
		}
	}
	return nil
}

func (m *DecoderService) MapPathParamsToStruct(
	//params map[string]string,
	r *http.Request,
	out any,
) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("output must be a non-nil pointer")
	}
	v = v.Elem()
	meta := m.GetStructMeta(v.Type())

	for tagVal, fi := range meta.FieldsByTag["path"] {
		//pathVal, ok := params[tagVal]
		pathVal := r.PathValue(tagVal)
		if pathVal == "" {
			continue
		}
		// if !ok || pathVal == "" {
		// 	continue
		// }
		field := v.Field(fi.Index)
		if err := fi.SetFunc(field, pathVal); err != nil {
			return fmt.Errorf("failed to set field %s: %w", fi.Name, err)
		}
	}
	return nil
}
