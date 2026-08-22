package ticketgate_test

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func TestBusinessChain37(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "business.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	svc := service.New(s, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		t.Fatal(err)
	}
	session, _, err := svc.NegotiateSession(fixture.GateID)
	if err != nil {
		t.Fatal(err)
	}
	record, err := svc.ValidateGateSession(fixture.RequestG34(session.ID))
	if err != nil {
		t.Fatal(err)
	}
	if record.Result != "validated" {
		t.Fatalf("expected complete validation record, got %q", record.Result)
	}
}
