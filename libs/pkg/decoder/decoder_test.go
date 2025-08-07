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
