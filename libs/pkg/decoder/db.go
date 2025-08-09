package decoder

import (
	"fmt"
	"reflect"
)

func (m *DecoderService) GetDbFieldsStruct(
	entity any,
) (map[string]string, error) {
	v := reflect.ValueOf(entity)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil, fmt.Errorf("entity must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	if cached, ok := m.dbFieldCache.Load(t); ok {
		return cached.(map[string]string), nil
	}

	meta := m.GetStructMeta(t)
	NameToTag := make(map[string]string)
	for tagVal, fi := range meta.FieldsByTag["db"] {
		NameToTag[fi.Name] = tagVal
	}

	m.dbFieldCache.Store(t, NameToTag)

	return NameToTag, nil
}

func (m *DecoderService) BuildUpdateMask(dto interface{}) []string {
	var mask []string
	if dto == nil {
		return mask
	}

	v := reflect.ValueOf(dto)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		// skip unexported fields
		if fieldType.PkgPath != "" {
			continue
		}
		fieldVal := v.Field(i)
		if fieldVal.Kind() == reflect.Pointer && !fieldVal.IsNil() {
			mask = append(mask, fieldType.Name)
		}
	}

	return mask
}
