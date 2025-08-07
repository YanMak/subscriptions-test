package decoder

import (
	"net/url"
	"testing"

	"github.com/google/uuid"
)

type queryIDs struct {
	IDs []uuid.UUID `query:"ids"`
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
