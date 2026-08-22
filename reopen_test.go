package ticketgate_test

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "reopen.db")
	s, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	svc := service.New(s, fixture.DefaultClock())
	if _, err := svc.RegisterGate(fixture.GateG34()); err != nil {
		t.Fatal(err)
	}
	if _, _, err := svc.NegotiateSession(fixture.GateID); err != nil {
		t.Fatal(err)
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
	s2, err := store.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()
	if _, err := s2.GetGate(fixture.GateID); err != nil {
		t.Fatal(err)
	}
	if _, err := s2.GetSession(fixture.SessionID()); err != nil {
		t.Fatal(err)
	}
}
