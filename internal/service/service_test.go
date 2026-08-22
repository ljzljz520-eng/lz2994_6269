package service

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/store"
)

func TestServiceRejectsDisabledGate(t *testing.T) {
	s, err := store.Open(filepath.Join(t.TempDir(), "svc.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	svc := New(s, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		t.Fatal(err)
	}
	if err := svc.DisableGate(fixture.GateID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := svc.NegotiateSession(fixture.GateID); err == nil {
		t.Fatal("expected disabled gate error")
	}
}
