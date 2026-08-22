package query_test

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/query"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func TestQueryManagement(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "query.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	svc := service.New(s, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := svc.NegotiateSession(fixture.GateID); err != nil {
		t.Fatal(err)
	}
	m := query.New(s)
	items, err := m.ListSessions(fixture.GateID, "active")
	if err != nil || len(items) != 1 {
		t.Fatalf("items=%d err=%v", len(items), err)
	}
	count, err := m.CountByResult(fixture.GateID, "negotiated")
	if err != nil || count != 1 {
		t.Fatalf("count=%d err=%v", count, err)
	}
}
