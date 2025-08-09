package dtodb

import (
	"reflect"
	"strings"
)

type TableNamer interface {
	TableName() string
}

type DtoToDbMapType map[string]string
type DtosToDbMapType map[string]DtoToDbMapType

type DtosToDbMapService struct {
	DtosToDbMap DtosToDbMapType
}

func NewDtosToDbMapService() *DtosToDbMapService {
	return &DtosToDbMapService{
		DtosToDbMap: make(DtosToDbMapType),
	}
}

func (m *DtosToDbMapService) AddMapping(entity TableNamer) error {
	mapping, err := DtoToDbMapping(entity)
	if err != nil {
		return err
	}
	m.DtosToDbMap[entity.TableName()] = mapping

	return nil
}

func (m *DtosToDbMapService) GetFieldMapping(tableName string, field string) string {
	res, ok := m.DtosToDbMap[tableName]
	if ok {
		return res[field]
	}
	return ""
}

func DtoToDbMapping(entityType any) (map[string]string, error) {

	m := make(map[string]string)

	tEntity := reflect.TypeOf(entityType).Elem()
	for j := 0; j < tEntity.NumField(); j++ {
		entityField := tEntity.Field(j)

		db := entityField.Tag.Get("db")
		if db == "" {
			continue
		}
		m[entityField.Name] = db
	}

	return m, nil
}

func BuildFieldmaskByJSONTags(updateDTO any, entityType any, exceptionStr []string) []string {
	mask := []string{}
	vDTO := reflect.ValueOf(updateDTO).Elem()
	tDTO := vDTO.Type()
	tEntity := reflect.TypeOf(entityType).Elem()
	for i := 0; i < tDTO.NumField(); i++ {
		dtoField := tDTO.Field(i)
		if dtoField.Name == "ID" || dtoField.Name == "UserID" {
			continue
		}
		val := vDTO.Field(i)
		if !val.IsNil() {
			jsonTag := dtoField.Tag.Get("json")
			if idx := strings.Index(jsonTag, ","); idx != -1 {
				jsonTag = jsonTag[:idx]
			}
			// look for entity with same field
			for j := 0; j < tEntity.NumField(); j++ {
				entityField := tEntity.Field(j)
				entityJsonTag := entityField.Tag.Get("json")
				if idx := strings.Index(entityJsonTag, ","); idx != -1 {
					entityJsonTag = entityJsonTag[:idx]
				}
				if jsonTag == entityJsonTag {
					mask = append(mask, entityField.Tag.Get("db"))
					break
				}
			}
		}
	}
	return mask
}
