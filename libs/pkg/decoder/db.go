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
		return nil, fmt.Errorf("output must be a non-nil pointer")
	}
	v = v.Elem()
	meta := m.GetStructMeta(v.Type())

	NameToTag := make(map[string]string)
	for tagVal, fi := range meta.FieldsByTag["path"] {
		//pathVal, ok := params[tagVal]
		NameToTag[fi.Name] = tagVal
	}
	return NameToTag, nil
}
