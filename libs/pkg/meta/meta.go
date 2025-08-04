package meta

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type FieldInfo struct {
	Name    string
	Index   int
	Type    reflect.Type
	Kind    reflect.Kind
	TagMap  map[string]string
	SetFunc func(reflect.Value, string) error
}

type StructMeta struct {
	FieldsByTag map[string]map[string]*FieldInfo // tagName -> tagValue -> FieldInfo
	Fields      []*FieldInfo
}

var typeCache sync.Map // reflect.Type -> *StructMeta

func GetStructMeta(t reflect.Type) *StructMeta {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if cached, ok := typeCache.Load(t); ok {
		return cached.(*StructMeta)
	}

	meta := &StructMeta{
		FieldsByTag: make(map[string]map[string]*FieldInfo),
		Fields:      make([]*FieldInfo, 0, t.NumField()),
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" {
			continue // skip unexported fields
		}

		tags := map[string]string{}
		for _, tagKey := range []string{"json", "db", "query", "path"} {
			val := field.Tag.Get(tagKey)
			val = strings.Split(val, ",")[0] // remove omitempty
			if val != "" {
				tags[tagKey] = val
			}
		}

		fi := &FieldInfo{
			Name:    field.Name,
			Index:   i,
			Type:    field.Type,
			Kind:    field.Type.Kind(),
			TagMap:  tags,
			SetFunc: defaultSetter(field.Type),
		}
		meta.Fields = append(meta.Fields, fi)

		for tagKey, tagVal := range tags {
			if meta.FieldsByTag[tagKey] == nil {
				meta.FieldsByTag[tagKey] = make(map[string]*FieldInfo)
			}
			meta.FieldsByTag[tagKey][tagVal] = fi
		}
	}
	typeCache.Store(t, meta)
	return meta
}

func MustGetFieldByTag(t reflect.Type, tagKey, tagValue string) (*FieldInfo, error) {
	meta := GetStructMeta(t)
	if fi, ok := meta.FieldsByTag[tagKey][tagValue]; ok {
		return fi, nil
	}
	return nil, fmt.Errorf("field with tag %s=\"%s\" not found in type %s", tagKey, tagValue, t.Name())
}

func defaultSetter(t reflect.Type) func(reflect.Value, string) error {
	if t.Kind() == reflect.Ptr {
		base := defaultSetter(t.Elem())
		return func(v reflect.Value, s string) error {
			ptr := reflect.New(t.Elem()).Elem()
			if err := base(ptr, s); err != nil {
				return err
			}
			v.Set(ptr.Addr())
			return nil
		}
	}

	switch t {
	case reflect.TypeOf(""):
		return func(v reflect.Value, s string) error {
			v.SetString(s)
			return nil
		}
	case reflect.TypeOf(0):
		return func(v reflect.Value, s string) error {
			val, err := parseInt(s)
			if err != nil {
				return err
			}
			v.SetInt(val)
			return nil
		}
	case reflect.TypeOf(true):
		return func(v reflect.Value, s string) error {
			b, err := parseBool(s)
			if err != nil {
				return err
			}
			v.SetBool(b)
			return nil
		}
	case reflect.TypeOf(uuid.UUID{}):
		return func(v reflect.Value, s string) error {
			u, err := uuid.Parse(s)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(u))
			return nil
		}
	case reflect.TypeOf(time.Time{}):
		return func(v reflect.Value, s string) error {
			t, err := time.Parse("2006-01-02", s)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(t))
			return nil
		}
	default:
		return func(v reflect.Value, s string) error {
			return fmt.Errorf("unsupported type: %v", t)
		}
	}
}

func parseInt(s string) (int64, error) {
	var i int64
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}

func parseBool(s string) (bool, error) {
	return s == "true" || s == "1", nil
}

func MapQueryToStruct(values url.Values, out any) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("output must be a non-nil pointer")
	}
	v = v.Elem()
	reflectedType := v.Type()
	meta := GetStructMeta(reflectedType)

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

func MapPathParamsToStruct(params map[string]string, out any) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("output must be a non-nil pointer")
	}
	v = v.Elem()
	meta := GetStructMeta(v.Type())

	for tagVal, fi := range meta.FieldsByTag["path"] {
		pathVal, ok := params[tagVal]
		if !ok || pathVal == "" {
			continue
		}
		field := v.Field(fi.Index)
		if err := fi.SetFunc(field, pathVal); err != nil {
			return fmt.Errorf("failed to set field %s: %w", fi.Name, err)
		}
	}
	return nil
}
