package ticketgate_test

import (
	"path/filepath"
	"testing"
	"ticketgate/internal/fixture"
	"ticketgate/internal/service"
	"ticketgate/internal/store"
)

func TestBadCiphertextAndParameterSet(t *testing.T) {
	path := filepath.Join(t.TempDir(), "errors.db")
	s, err := store.Open(path)
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
	session.Ciphertext = fixture.InvalidCiphertext()
	if err := s.SaveSession(session); err != nil {
		t.Fatal(err)
	}
	if err := svc.ValidateCiphertext(fixture.RequestG34(session.ID)); err == nil {
		t.Fatal("expected ciphertext error")
	}
	if _, err := svc.ValidateGateSession(fixture.WrongParameterRequest(session.ID)); err == nil {
		t.Fatal("expected parameter error")
	}
}

func TestMissingPersistenceFile(t *testing.T) {
	if _, err := store.Open(t.TempDir()); err == nil {
		t.Fatal("expected directory open error")
	}
}
