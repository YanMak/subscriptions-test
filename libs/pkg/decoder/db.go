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
	meta := m.GetStructMeta(v.Type())

	NameToTag := make(map[string]string)
	for tagVal, fi := range meta.FieldsByTag["db"] {
		//pathVal, ok := params[tagVal]
		NameToTag[fi.Name] = tagVal
	}
	return NameToTag, nil
}
