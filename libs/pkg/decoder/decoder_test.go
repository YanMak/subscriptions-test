package decoder

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/google/uuid"
	customerHttpContract "project1.v0/contracts/transport/http/customer"
	validatorx "project1.v0/pkg/validator_x"
)

type queryIDs struct {
	IDs []uuid.UUID `query:"ids"`
}

type priceRange struct {
	PriceFrom int `db:"price"`
	PriceTo   int `db:"price"`
}

func TestDefaultSetterUUIDSlice(t *testing.T) {
	id1 := uuid.New()
	id2 := uuid.New()
	values := url.Values{}
	values.Set("ids", id1.String()+","+id2.String())

	var q queryIDs
	d := NewDecoderService(DecoderServiceDeps{})
	if err := d.MapQueryToStruct(values, &q); err != nil {
		t.Fatalf("MapQueryToStruct returned error: %v", err)
	}

	if len(q.IDs) != 2 {
		t.Fatalf("expected 2 ids, got %d", len(q.IDs))
	}
	if q.IDs[0] != id1 || q.IDs[1] != id2 {
		t.Fatalf("unexpected ids slice: %v", q.IDs)
	}
}

type querySort struct {
	Sort []struct {
		Field     string
		Direction string
	} `query:"sort"`
}

func TestDefaultSetterSortSlice(t *testing.T) {
	values := url.Values{}
	values.Set("sort", "price:asc,start_date:desc")

	var q querySort
	d := NewDecoderService(DecoderServiceDeps{})
	if err := d.MapQueryToStruct(values, &q); err != nil {
		t.Fatalf("MapQueryToStruct returned error: %v", err)
	}

	if len(q.Sort) != 2 {
		t.Fatalf("expected 2 sort items, got %d", len(q.Sort))
	}
	if q.Sort[0].Field != "price" || q.Sort[0].Direction != "asc" {
		t.Fatalf("unexpected first sort item: %#v", q.Sort[0])
	}
	if q.Sort[1].Field != "start_date" || q.Sort[1].Direction != "desc" {
		t.Fatalf("unexpected second sort item: %#v", q.Sort[1])
	}
}

func TestDefaultSetterSortSliceInvalidDirection(t *testing.T) {
	values := url.Values{}
	values.Set("sort", "price:up")

	var q querySort
	d := NewDecoderService(DecoderServiceDeps{})
	if err := d.MapQueryToStruct(values, &q); err == nil {
		t.Fatalf("expected error for invalid direction, got nil")
	}
}

func TestSubscriptionListQueryInvalidField(t *testing.T) {
	values := url.Values{}
	values.Set("sort", "unknown:asc")

	var q customerHttpContract.SubscriptionListQuery
	d := NewDecoderService(DecoderServiceDeps{})
	if err := d.MapQueryToStruct(values, &q); err != nil {
		t.Fatalf("MapQueryToStruct returned error: %v", err)
	}

	v := validatorx.NewValidatorX()
	if err := v.Struct(&q); err == nil {
		t.Fatalf("expected validation error for invalid field, got nil")
	}
}

func TestGetStructMetaDuplicateTags(t *testing.T) {
	d := NewDecoderService(DecoderServiceDeps{})

	meta := d.GetStructMeta(reflect.TypeOf(priceRange{}))
	if len(meta.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(meta.Fields))
	}
	fields := meta.FieldsByTag["db"]["price"]
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields for tag price, got %d", len(fields))
	}
}
